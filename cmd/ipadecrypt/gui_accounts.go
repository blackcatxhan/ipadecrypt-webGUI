package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/londek/ipadecrypt/internal/appstore"
	"github.com/londek/ipadecrypt/internal/config"
)

const (
	accountsFileName = "accounts.json"
	legacyAccountID  = "legacy"
)

// AccountRecord represents one saved Apple ID account metadata and tokens.
type AccountRecord struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password,omitempty"`
	PasswordToken string `json:"passwordToken,omitempty"`
	DSID          string `json:"directoryServicesIdentifier,omitempty"`
	StoreFront    string `json:"storeFront,omitempty"`
	Pod           string `json:"pod,omitempty"`
	Name          string `json:"name,omitempty"`
	CreatedAt     int64  `json:"createdAt"`
	LastLoginAt   int64  `json:"lastLoginAt,omitempty"`
}

// AccountsFile represents the JSON structure saved to ~/.ipadecrypt/accounts.json.
type AccountsFile struct {
	Version  int             `json:"version"`
	ActiveID string          `json:"activeAccountId,omitempty"`
	Accounts []AccountRecord `json:"accounts"`
}

// AccountView is the safe public representation sent to the Web GUI.
type AccountView struct {
	ID                string `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name,omitempty"`
	Storefront        string `json:"storefront,omitempty"`
	StorefrontCountry string `json:"storefrontCountry,omitempty"`
	DSID              string `json:"dsid,omitempty"`
	Authenticated     bool   `json:"authenticated"`
	Active            bool   `json:"active"`
	HasCredential     bool   `json:"hasCredential"`
	CredentialFile    string `json:"credentialFile,omitempty"`
}

// PendingLogin tracks a 2FA challenge in progress.
type PendingLogin struct {
	Email     string
	Password  string
	AccountID string // if re-authenticating existing account
	CreatedAt time.Time
}

// AccountManager coordinates multi-account persistence, sessions, and switching.
type AccountManager struct {
	mu            sync.RWMutex
	paths         *config.Paths
	cfg           *config.Config
	accounts      map[string]*AccountRecord
	activeID      string
	clients       map[string]*appstore.Client
	pendingLogins map[string]*PendingLogin // keyed by normalized email
}

var accountsManager *AccountManager

func newAccountID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func initAccountManager(cfg *config.Config, paths *config.Paths) (*AccountManager, error) {
	mgr := &AccountManager{
		paths:         paths,
		cfg:           cfg,
		accounts:      make(map[string]*AccountRecord),
		clients:       make(map[string]*appstore.Client),
		pendingLogins: make(map[string]*PendingLogin),
	}

	accFile := filepath.Join(paths.Root, accountsFileName)
	data, err := os.ReadFile(accFile)
	if err == nil {
		var f AccountsFile
		if err := json.Unmarshal(data, &f); err == nil {
			for i := range f.Accounts {
				rec := f.Accounts[i]
				mgr.accounts[rec.ID] = &rec
			}
			mgr.activeID = f.ActiveID
		}
	}

	// Migrate legacy account in config.Apple if accounts.json was empty
	if len(mgr.accounts) == 0 && cfg.Apple.Email != "" {
		id := legacyAccountID
		rec := &AccountRecord{
			ID:            id,
			Email:         cfg.Apple.Email,
			Password:      cfg.Apple.Password,
			PasswordToken: cfg.Apple.PasswordToken,
			DSID:          cfg.Apple.DirectoryServicesIdentifier,
			StoreFront:    cfg.Apple.StoreFront,
			Pod:           cfg.Apple.Pod,
			CreatedAt:     time.Now().Unix(),
		}
		mgr.accounts[id] = rec
		mgr.activeID = id
		_ = mgr.save()
	}

	// If activeID is invalid or empty, pick the first available
	if mgr.activeID == "" || mgr.accounts[mgr.activeID] == nil {
		for id := range mgr.accounts {
			mgr.activeID = id
			break
		}
		if mgr.activeID != "" {
			_ = mgr.save()
		}
	}

	// Keep config.Apple in sync with active account
	mgr.syncConfigApple()

	accountsManager = mgr
	return mgr, nil
}

func (m *AccountManager) save() error {
	f := AccountsFile{
		Version:  1,
		ActiveID: m.activeID,
		Accounts: make([]AccountRecord, 0, len(m.accounts)),
	}
	for _, rec := range m.accounts {
		if rec != nil {
			f.Accounts = append(f.Accounts, *rec)
		}
	}

	// Sort accounts by email for deterministic output
	sort.Slice(f.Accounts, func(i, j int) bool {
		return strings.ToLower(f.Accounts[i].Email) < strings.ToLower(f.Accounts[j].Email)
	})

	data, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	accFile := filepath.Join(m.paths.Root, accountsFileName)
	return os.WriteFile(accFile, data, 0o600)
}

func (m *AccountManager) syncConfigApple() {
	if m.activeID == "" {
		return
	}
	active := m.accounts[m.activeID]
	if active == nil {
		return
	}
	m.cfg.Apple.Email = active.Email
	m.cfg.Apple.Password = active.Password
	m.cfg.Apple.PasswordToken = active.PasswordToken
	m.cfg.Apple.DirectoryServicesIdentifier = active.DSID
	m.cfg.Apple.StoreFront = active.StoreFront
	m.cfg.Apple.Pod = active.Pod
	_ = m.cfg.Save()
}

func (m *AccountManager) getCookiesPath(accountID string) string {
	if accountID == "" || accountID == legacyAccountID {
		return filepath.Join(m.paths.Root, "cookies")
	}
	dir := filepath.Join(m.paths.Root, "accounts", accountID)
	_ = os.MkdirAll(dir, 0o755)
	return filepath.Join(dir, "cookies")
}

// ClientForAccount returns an appstore.Client configured with this account's cookie store.
func (m *AccountManager) ClientForAccount(accountID string) (*appstore.Client, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if accountID == "" {
		accountID = m.activeID
	}
	if client, ok := m.clients[accountID]; ok {
		return client, nil
	}

	cookiesPath := m.getCookiesPath(accountID)
	client, err := appstore.New(cookiesPath)
	if err != nil {
		return nil, fmt.Errorf("create appstore client for %s: %w", accountID, err)
	}

	m.clients[accountID] = client
	return client, nil
}

// ActiveAccount returns the currently active AccountRecord.
func (m *AccountManager) ActiveAccount() (*AccountRecord, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.activeID == "" {
		return nil, false
	}
	rec, ok := m.accounts[m.activeID]
	return rec, ok
}

// GetAccount returns an account by ID, or the active account if id is empty.
func (m *AccountManager) GetAccount(id string) (*AccountRecord, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if id == "" {
		id = m.activeID
	}
	rec, ok := m.accounts[id]
	return rec, ok
}

// AsAppStoreAccount converts an AccountRecord to *appstore.Account.
func (r *AccountRecord) AsAppStoreAccount() *appstore.Account {
	if r == nil {
		return nil
	}
	return &appstore.Account{
		Email:               r.Email,
		Password:            r.Password,
		PasswordToken:       r.PasswordToken,
		DirectoryServicesID: r.DSID,
		StoreFront:          r.StoreFront,
		Pod:                 r.Pod,
		Name:                r.Name,
	}
}

// Views returns safe account view objects for the Web GUI.
func (m *AccountManager) Views() ([]AccountView, string) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	views := make([]AccountView, 0, len(m.accounts))
	for _, rec := range m.accounts {
		country := ""
		if rec.StoreFront != "" {
			if cc, err := appstore.CountryCodeFromStoreFront(rec.StoreFront); err == nil {
				country = cc
			}
		}

		hasCred := false
		credPath := ""
		if rec.DSID != "" {
			hasCred = credentialCached(m.paths, rec.DSID)
			if p, err := m.paths.CredentialPath(rec.DSID); err == nil {
				credPath = p
			}
		}

		views = append(views, AccountView{
			ID:                rec.ID,
			Email:             rec.Email,
			Name:              rec.Name,
			Storefront:        rec.StoreFront,
			StorefrontCountry: country,
			DSID:              rec.DSID,
			Authenticated:     rec.PasswordToken != "",
			Active:            rec.ID == m.activeID,
			HasCredential:     hasCred,
			CredentialFile:    credPath,
		})
	}

	sort.Slice(views, func(i, j int) bool {
		if views[i].Active != views[j].Active {
			return views[i].Active
		}
		return strings.ToLower(views[i].Email) < strings.ToLower(views[j].Email)
	})

	return views, m.activeID
}

// SetActive switches the active account and keeps config.json in sync.
func (m *AccountManager) SetActive(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	rec, ok := m.accounts[id]
	if !ok || rec == nil {
		return fmt.Errorf("account %q not found", id)
	}

	m.activeID = id
	m.syncConfigApple()
	return m.save()
}

// RemoveAccount deletes an account from the saved list.
func (m *AccountManager) RemoveAccount(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.accounts[id]; !ok {
		return fmt.Errorf("account %q not found", id)
	}

	delete(m.accounts, id)
	delete(m.clients, id)

	if m.activeID == id {
		m.activeID = ""
		for nextID := range m.accounts {
			m.activeID = nextID
			break
		}
	}

	m.syncConfigApple()
	return m.save()
}

// Login attempts to authenticate against the App Store.
// If 2FA is needed, it saves pending state and returns (nil, true, nil).
func (m *AccountManager) Login(email, password, authCode string) (*AccountRecord, bool, error) {
	email = strings.TrimSpace(email)
	normEmail := strings.ToLower(email)
	if email == "" {
		return nil, false, errors.New("email is required")
	}

	// Check if this email is an existing account to update
	var targetID string
	m.mu.RLock()
	for id, rec := range m.accounts {
		if strings.EqualFold(rec.Email, email) {
			targetID = id
			break
		}
	}
	m.mu.RUnlock()

	if targetID == "" {
		targetID = newAccountID()
	}

	client, err := m.ClientForAccount(targetID)
	if err != nil {
		return nil, false, err
	}

	// Try login
	acc, err := client.Login(email, password, authCode)
	if err != nil {
		if errors.Is(err, appstore.ErrAuthCodeRequired) {
			m.mu.Lock()
			m.pendingLogins[normEmail] = &PendingLogin{
				Email:     email,
				Password:  password,
				AccountID: targetID,
				CreatedAt: time.Now(),
			}
			m.mu.Unlock()
			return nil, true, nil // 2FA code required
		}
		return nil, false, err
	}

	// Success! Save or update AccountRecord
	m.mu.Lock()
	delete(m.pendingLogins, normEmail)

	rec := &AccountRecord{
		ID:            targetID,
		Email:         acc.Email,
		Password:      password,
		PasswordToken: acc.PasswordToken,
		DSID:          acc.DirectoryServicesID,
		StoreFront:    acc.StoreFront,
		Pod:           acc.Pod,
		Name:          acc.Name,
		LastLoginAt:   time.Now().Unix(),
	}
	if existing, ok := m.accounts[targetID]; ok && existing != nil {
		rec.CreatedAt = existing.CreatedAt
	} else {
		rec.CreatedAt = time.Now().Unix()
	}

	m.accounts[targetID] = rec
	m.activeID = targetID
	m.syncConfigApple()
	_ = m.save()
	m.mu.Unlock()

	return rec, false, nil
}

// Reauth refreshes an account's token if expired.
func (m *AccountManager) Reauth(accountID string) (*AccountRecord, error) {
	m.mu.RLock()
	rec, ok := m.accounts[accountID]
	m.mu.RUnlock()
	if !ok || rec == nil {
		return nil, errors.New("account not found")
	}
	if rec.Password == "" {
		return nil, errors.New("no stored password for account; please re-login")
	}

	updated, _, err := m.Login(rec.Email, rec.Password, "")
	return updated, err
}

// HTTP API Handlers for accounts

func handleAPIAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	views, activeID := accountsManager.Views()
	jsonResponse(w, http.StatusOK, map[string]any{
		"success":  true,
		"activeId": activeID,
		"accounts": views,
	})
}

func handleAPIActivateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var req struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ID == "" {
		jsonError(w, http.StatusBadRequest, "account id is required")
		return
	}
	if err := accountsManager.SetActive(req.ID); err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]any{
		"success":  true,
		"activeId": req.ID,
	})
}

func handleAPIRemoveAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var req struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ID == "" {
		jsonError(w, http.StatusBadRequest, "account id is required")
		return
	}
	if err := accountsManager.RemoveAccount(req.ID); err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	views, activeID := accountsManager.Views()
	jsonResponse(w, http.StatusOK, map[string]any{
		"success":  true,
		"activeId": activeID,
		"accounts": views,
	})
}

func handleAPILogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		AuthCode string `json:"authCode"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	normEmail := strings.ToLower(strings.TrimSpace(req.Email))
	// If 2FA code is provided without password, check pendingLogins
	if req.AuthCode != "" && req.Password == "" {
		accountsManager.mu.RLock()
		pending := accountsManager.pendingLogins[normEmail]
		accountsManager.mu.RUnlock()
		if pending != nil {
			req.Password = pending.Password
		}
	}

	rec, authCodeRequired, err := accountsManager.Login(req.Email, req.Password, req.AuthCode)
	if err != nil {
		jsonResponse(w, http.StatusOK, map[string]any{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if authCodeRequired {
		jsonResponse(w, http.StatusOK, map[string]any{
			"success":          false,
			"authCodeRequired": true,
			"message":          "Vui lòng nhập mã xác thực 2 bước (2FA) gồm 6 chữ số gửi về thiết bị Apple của bạn.",
		})
		return
	}

	country := ""
	if rec.StoreFront != "" {
		if cc, err := appstore.CountryCodeFromStoreFront(rec.StoreFront); err == nil {
			country = cc
		}
	}

	hasCred := false
	if rec.DSID != "" {
		hasCred = credentialCached(accountsManager.paths, rec.DSID)
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"account": AccountView{
			ID:                rec.ID,
			Email:             rec.Email,
			Name:              rec.Name,
			Storefront:        rec.StoreFront,
			StorefrontCountry: country,
			DSID:              rec.DSID,
			Authenticated:     true,
			Active:            true,
			HasCredential:     hasCred,
		},
	})
}
