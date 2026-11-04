package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type BookmarkedApp struct {
	ID                   int64   `json:"id,omitempty"` // trackId
	BundleID             string  `json:"bundleId"`
	Name                 string  `json:"name"`
	Developer            string  `json:"developer,omitempty"`
	IconURL              string  `json:"iconUrl,omitempty"`
	Country              string  `json:"country,omitempty"`
	Price                float64 `json:"price,omitempty"`
	FormattedPrice       string  `json:"formattedPrice,omitempty"`
	IsFree               bool    `json:"isFree"`
	AddedAt              int64   `json:"addedAt"`
	LatestStoreVersion   string  `json:"latestStoreVersion,omitempty"`
	LatestReleaseDate    string  `json:"latestReleaseDate,omitempty"`
	LastDecryptedVersion string  `json:"lastDecryptedVersion,omitempty"`
	DecryptedFilePath    string  `json:"decryptedFilePath,omitempty"`
	HasUpdate            bool    `json:"hasUpdate"`
	LastCheckedAt        int64   `json:"lastCheckedAt,omitempty"`
}

type BookmarksStore struct {
	mu        sync.RWMutex
	filePath  string
	Bookmarks map[string]*BookmarkedApp `json:"bookmarks"` // bundleId -> BookmarkedApp
}

var globalBookmarks *BookmarksStore

func initBookmarksStore(rootDir string) (*BookmarksStore, error) {
	fp := filepath.Join(rootDir, "bookmarks.json")
	store := &BookmarksStore{
		filePath:  fp,
		Bookmarks: make(map[string]*BookmarkedApp),
	}

	data, err := os.ReadFile(fp)
	if err == nil {
		var list []*BookmarkedApp
		// Try unmarshaling either a map or an array for backwards/clean format
		if err := json.Unmarshal(data, &list); err == nil {
			for _, item := range list {
				if item.BundleID != "" {
					store.Bookmarks[item.BundleID] = item
				}
			}
		} else {
			var m map[string]*BookmarkedApp
			if err := json.Unmarshal(data, &m); err == nil {
				store.Bookmarks = m
			}
		}
	}

	globalBookmarks = store
	return store, nil
}

func (s *BookmarksStore) save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]*BookmarkedApp, 0, len(s.Bookmarks))
	for _, v := range s.Bookmarks {
		list = append(list, v)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].AddedAt > list[j].AddedAt
	})

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	if s.filePath == "" {
		s.filePath = filepath.Join(".", "bookmarks.json")
	}

	_ = os.MkdirAll(filepath.Dir(s.filePath), 0o755)
	return os.WriteFile(s.filePath, data, 0o600)
}

// syncWithDecryptedFiles checks local decrypted folder and updates lastDecryptedVersion
func (s *BookmarksStore) syncWithDecryptedFiles() {
	files, err := listDecryptedFiles()
	if err != nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, b := range s.Bookmarks {
		var latestDecVer string
		var latestDecPath string
		var latestModTime int64

		for _, f := range files {
			if f.BundleID == b.BundleID || strings.HasPrefix(strings.ToLower(f.Name), strings.ToLower(b.BundleID)+"_") {
				if f.ModifiedAt > latestModTime {
					latestModTime = f.ModifiedAt
					latestDecVer = f.Version
					latestDecPath = f.Path
				}
			}
		}

		if latestDecVer != "" {
			b.LastDecryptedVersion = latestDecVer
			b.DecryptedFilePath = latestDecPath
		} else if b.DecryptedFilePath != "" {
			if _, err := os.Stat(b.DecryptedFilePath); os.IsNotExist(err) {
				b.LastDecryptedVersion = ""
				b.DecryptedFilePath = ""
			}
		}

		if b.LatestStoreVersion != "" && b.LastDecryptedVersion != "" {
			b.HasUpdate = isNewerVersion(b.LatestStoreVersion, b.LastDecryptedVersion)
		} else {
			b.HasUpdate = false
		}
	}
}

// isNewerVersion compares two version strings (e.g. "8.9.24" vs "8.9.20")
func isNewerVersion(storeVer, decVer string) bool {
	storeVer = strings.TrimSpace(storeVer)
	decVer = strings.TrimSpace(decVer)
	if storeVer == "" || decVer == "" {
		return false
	}
	if storeVer == decVer {
		return false
	}

	// Compare component by component
	storeParts := strings.Split(storeVer, ".")
	decParts := strings.Split(decVer, ".")

	maxLen := len(storeParts)
	if len(decParts) > maxLen {
		maxLen = len(decParts)
	}

	for i := 0; i < maxLen; i++ {
		var sNum, dNum int64
		if i < len(storeParts) {
			sNum, _ = strconv.ParseInt(cleanVersionPart(storeParts[i]), 10, 64)
		}
		if i < len(decParts) {
			dNum, _ = strconv.ParseInt(cleanVersionPart(decParts[i]), 10, 64)
		}

		if sNum > dNum {
			return true
		} else if sNum < dNum {
			return false
		}
	}

	// Fallback to string inequality if numeric parts match but strings differ
	return storeVer != decVer
}

func cleanVersionPart(p string) string {
	var sb strings.Builder
	for _, c := range p {
		if c >= '0' && c <= '9' {
			sb.WriteRune(c)
		} else {
			break
		}
	}
	return sb.String()
}

func fetchStoreMetadata(bundleID string, trackID int64, country string) (version string, releaseDate string, name string, dev string, icon string, price float64, formattedPrice string, currency string, err error) {
	if country == "" {
		country = "US"
	}
	q := url.Values{}
	q.Set("country", country)
	q.Set("entity", "software,iPadSoftware")
	if trackID > 0 {
		q.Set("id", fmt.Sprintf("%d", trackID))
	} else {
		q.Set("bundleId", bundleID)
	}

	reqURL := fmt.Sprintf("https://itunes.apple.com/lookup?%s", q.Encode())
	client := &http.Client{Timeout: 8 * time.Second}
	resp, err := client.Get(reqURL)
	if err != nil {
		return "", "", "", "", "", 0, "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", "", "", "", 0, "", "", fmt.Errorf("status %d", resp.StatusCode)
	}

	var res struct {
		Results []struct {
			TrackID                   int64   `json:"trackId"`
			TrackName                 string  `json:"trackName"`
			BundleID                  string  `json:"bundleId"`
			Version                   string  `json:"version"`
			CurrentVersionReleaseDate string  `json:"currentVersionReleaseDate"`
			ArtistName                string  `json:"artistName"`
			ArtworkUrl512             string  `json:"artworkUrl512"`
			ArtworkUrl100             string  `json:"artworkUrl100"`
			Price                     float64 `json:"price"`
			FormattedPrice            string  `json:"formattedPrice"`
			Currency                  string  `json:"currency"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", "", "", "", "", 0, "", "", err
	}

	if len(res.Results) == 0 {
		return "", "", "", "", "", 0, "", "", fmt.Errorf("app not found on store")
	}

	item := res.Results[0]
	iconUrl := item.ArtworkUrl512
	if iconUrl == "" {
		iconUrl = item.ArtworkUrl100
	}

	return item.Version, item.CurrentVersionReleaseDate, item.TrackName, item.ArtistName, iconUrl, item.Price, item.FormattedPrice, item.Currency, nil
}

// Handlers

func handleAPIBookmarks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	if globalBookmarks == nil {
		jsonResponse(w, http.StatusOK, map[string]any{
			"success":   true,
			"bookmarks": []*BookmarkedApp{},
			"count":     0,
			"hasUpdate": 0,
		})
		return
	}

	globalBookmarks.syncWithDecryptedFiles()

	globalBookmarks.mu.RLock()
	list := make([]*BookmarkedApp, 0, len(globalBookmarks.Bookmarks))
	updateCount := 0
	for _, b := range globalBookmarks.Bookmarks {
		list = append(list, b)
		if b.HasUpdate {
			updateCount++
		}
	}
	globalBookmarks.mu.RUnlock()

	sort.Slice(list, func(i, j int) bool {
		if list[i].HasUpdate != list[j].HasUpdate {
			return list[i].HasUpdate // items with updates first
		}
		return list[i].AddedAt > list[j].AddedAt
	})

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":      true,
		"bookmarks":    list,
		"count":        len(list),
		"updatesCount": updateCount,
	})
}

func handleAPIBookmarksToggle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req BookmarkedApp
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	req.BundleID = strings.TrimSpace(req.BundleID)
	if req.BundleID == "" {
		jsonError(w, http.StatusBadRequest, "bundleId is required")
		return
	}

	if globalBookmarks == nil {
		jsonError(w, http.StatusInternalServerError, "bookmarks store not initialized")
		return
	}

	globalBookmarks.mu.Lock()
	_, exists := globalBookmarks.Bookmarks[req.BundleID]
	if exists {
		delete(globalBookmarks.Bookmarks, req.BundleID)
		globalBookmarks.mu.Unlock()
		_ = globalBookmarks.save()
		jsonResponse(w, http.StatusOK, map[string]any{
			"success":    true,
			"bookmarked": false,
			"bundleId":   req.BundleID,
			"message":    "Đã bỏ theo dõi ứng dụng",
		})
		return
	}

	// Add new bookmark
	app := &BookmarkedApp{
		ID:             req.ID,
		BundleID:       req.BundleID,
		Name:           req.Name,
		Developer:      req.Developer,
		IconURL:        req.IconURL,
		Country:        req.Country,
		Price:          req.Price,
		FormattedPrice: req.FormattedPrice,
		IsFree:               req.IsFree,
		LastDecryptedVersion: req.LastDecryptedVersion,
		DecryptedFilePath:    req.DecryptedFilePath,
		AddedAt:              time.Now().Unix(),
		LastCheckedAt:  time.Now().Unix(),
	}
	globalBookmarks.Bookmarks[req.BundleID] = app
	globalBookmarks.mu.Unlock()

	// Sync local decrypted version immediately
	globalBookmarks.syncWithDecryptedFiles()

	// Query App Store for latest version in background/goroutine so save is fast
	go func(b *BookmarkedApp) {
		ver, relDate, name, dev, icon, _, _, _, err := fetchStoreMetadata(b.BundleID, b.ID, b.Country)
		if err == nil {
			globalBookmarks.mu.Lock()
			b.LatestStoreVersion = ver
			b.LatestReleaseDate = relDate
			if b.Name == "" {
				b.Name = name
			}
			if b.Developer == "" {
				b.Developer = dev
			}
			if b.IconURL == "" {
				b.IconURL = icon
			}
			b.LastCheckedAt = time.Now().Unix()
			if b.LatestStoreVersion != "" && b.LastDecryptedVersion != "" {
				b.HasUpdate = isNewerVersion(b.LatestStoreVersion, b.LastDecryptedVersion)
			}
			globalBookmarks.mu.Unlock()
			_ = globalBookmarks.save()
		}
	}(app)

	_ = globalBookmarks.save()

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":    true,
		"bookmarked": true,
		"app":        app,
		"bundleId":   req.BundleID,
		"message":    "Đã thêm vào danh sách theo dõi!",
	})
}

func handleAPIBookmarksRemove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		BundleID string `json:"bundleId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	req.BundleID = strings.TrimSpace(req.BundleID)
	if req.BundleID == "" {
		jsonError(w, http.StatusBadRequest, "bundleId is required")
		return
	}

	if globalBookmarks == nil {
		jsonError(w, http.StatusInternalServerError, "bookmarks store not initialized")
		return
	}

	globalBookmarks.mu.Lock()
	delete(globalBookmarks.Bookmarks, req.BundleID)
	globalBookmarks.mu.Unlock()
	_ = globalBookmarks.save()

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":  true,
		"bundleId": req.BundleID,
		"message":  "Đã xóa khỏi danh sách theo dõi",
	})
}

func handleAPIBookmarksCheckUpdates(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	if globalBookmarks == nil {
		jsonError(w, http.StatusInternalServerError, "bookmarks store not initialized")
		return
	}

	globalBookmarks.syncWithDecryptedFiles()

	globalBookmarks.mu.RLock()
	apps := make([]*BookmarkedApp, 0, len(globalBookmarks.Bookmarks))
	for _, b := range globalBookmarks.Bookmarks {
		apps = append(apps, b)
	}
	globalBookmarks.mu.RUnlock()

	// Query each app's latest store version
	for _, app := range apps {
		ver, relDate, name, dev, icon, price, fmtPrice, _, err := fetchStoreMetadata(app.BundleID, app.ID, app.Country)
		if err == nil {
			globalBookmarks.mu.Lock()
			app.LatestStoreVersion = ver
			app.LatestReleaseDate = relDate
			if app.Name == "" {
				app.Name = name
			}
			if app.Developer == "" {
				app.Developer = dev
			}
			if app.IconURL == "" {
				app.IconURL = icon
			}
			if app.Price == 0 && price > 0 {
				app.Price = price
				app.FormattedPrice = fmtPrice
			}
			app.LastCheckedAt = time.Now().Unix()

			if app.LatestStoreVersion != "" && app.LastDecryptedVersion != "" {
				app.HasUpdate = isNewerVersion(app.LatestStoreVersion, app.LastDecryptedVersion)
			} else {
				app.HasUpdate = false
			}
			globalBookmarks.mu.Unlock()
		}
		time.Sleep(150 * time.Millisecond) // gentle pace
	}

	_ = globalBookmarks.save()

	// Return updated list
	globalBookmarks.mu.RLock()
	list := make([]*BookmarkedApp, 0, len(globalBookmarks.Bookmarks))
	updatesCount := 0
	for _, b := range globalBookmarks.Bookmarks {
		list = append(list, b)
		if b.HasUpdate {
			updatesCount++
		}
	}
	globalBookmarks.mu.RUnlock()

	sort.Slice(list, func(i, j int) bool {
		if list[i].HasUpdate != list[j].HasUpdate {
			return list[i].HasUpdate
		}
		return list[i].AddedAt > list[j].AddedAt
	})

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":      true,
		"bookmarks":    list,
		"updatesCount": updatesCount,
		"message":      fmt.Sprintf("Đã kiểm tra xong. Phát hiện %d ứng dụng có bản cập nhật mới!", updatesCount),
	})
}
