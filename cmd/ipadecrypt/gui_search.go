package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/londek/ipadecrypt/internal/appstore"
)

type SearchAppResult struct {
	ID               int64    `json:"id"`
	BundleID         string   `json:"bundleId"`
	Name             string   `json:"name"`
	Version          string   `json:"version"`
	Price            float64  `json:"price"`
	FormattedPrice   string   `json:"formattedPrice"`
	Currency         string   `json:"currency,omitempty"`
	IsFree           bool     `json:"isFree"`
	IconURL          string   `json:"iconUrl"`
	Developer        string   `json:"developer"`
	Description      string   `json:"description,omitempty"`
	MinimumOSVersion string   `json:"minimumOsVersion,omitempty"`
	Genre            string   `json:"genre,omitempty"`
	SupportedDevices []string `json:"supportedDevices,omitempty"`
	Platforms        []string `json:"platforms,omitempty"`
}

type itunesSearchResponse struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		TrackID          int64    `json:"trackId"`
		TrackName        string   `json:"trackName"`
		BundleID         string   `json:"bundleId"`
		Version          string   `json:"version"`
		Price            float64  `json:"price"`
		FormattedPrice   string   `json:"formattedPrice"`
		Currency         string   `json:"currency"`
		ArtworkUrl100    string   `json:"artworkUrl100"`
		ArtworkUrl512    string   `json:"artworkUrl512"`
		ArtistName       string   `json:"artistName"`
		Description      string   `json:"description"`
		MinimumOSVersion string   `json:"minimumOsVersion"`
		PrimaryGenreName string   `json:"primaryGenreName"`
		SupportedDevices []string `json:"supportedDevices"`
		Kind             string   `json:"kind"`
	} `json:"results"`
}

func searchAppStore(term string, country string, limit int) ([]SearchAppResult, error) {
	term = strings.TrimSpace(term)
	if term == "" {
		return nil, nil
	}
	if limit <= 0 || limit > 50 {
		limit = 25
	}
	if country == "" {
		country = "US"
	}

	params := url.Values{}
	params.Set("term", term)
	params.Set("country", country)
	params.Set("entity", "software,iPadSoftware")
	params.Set("media", "software")
	params.Set("limit", strconv.Itoa(limit))

	searchURL := fmt.Sprintf("https://itunes.apple.com/search?%s", params.Encode())

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "iTunes/12.11.3 (Windows; Microsoft Windows 10 x64)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("iTunes search returned HTTP %d", resp.StatusCode)
	}

	var searchResp itunesSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("decode search response: %w", err)
	}

	out := make([]SearchAppResult, 0, len(searchResp.Results))
	for _, item := range searchResp.Results {
		if item.BundleID == "" && item.TrackID == 0 {
			continue
		}

		icon := item.ArtworkUrl512
		if icon == "" {
			icon = item.ArtworkUrl100
		}

		platforms := make([]string, 0)
		for _, dev := range item.SupportedDevices {
			if strings.HasPrefix(dev, "iPhone") && !containsStr(platforms, "iOS") {
				platforms = append(platforms, "iOS")
			} else if strings.HasPrefix(dev, "iPad") && !containsStr(platforms, "iPadOS") {
				platforms = append(platforms, "iPadOS")
			} else if strings.HasPrefix(dev, "Watch") && !containsStr(platforms, "watchOS") {
				platforms = append(platforms, "watchOS")
			}
		}
		if len(platforms) == 0 {
			platforms = append(platforms, "iOS")
		}

		isFree := item.Price == 0 || strings.EqualFold(item.FormattedPrice, "Free") || strings.EqualFold(item.FormattedPrice, "Miễn phí")
		formattedPrice := item.FormattedPrice
		if formattedPrice == "" {
			if isFree {
				formattedPrice = "Free"
			} else if item.Price > 0 {
				if item.Currency != "" {
					formattedPrice = fmt.Sprintf("%.2f %s", item.Price, item.Currency)
				} else {
					formattedPrice = fmt.Sprintf("$%.2f", item.Price)
				}
			}
		}

		out = append(out, SearchAppResult{
			ID:               item.TrackID,
			BundleID:         item.BundleID,
			Name:             item.TrackName,
			Version:          item.Version,
			Price:            item.Price,
			FormattedPrice:   formattedPrice,
			Currency:         item.Currency,
			IsFree:           isFree,
			IconURL:          icon,
			Developer:        item.ArtistName,
			Description:      item.Description,
			MinimumOSVersion: item.MinimumOSVersion,
			Genre:            item.PrimaryGenreName,
			SupportedDevices: item.SupportedDevices,
			Platforms:        platforms,
		})
	}

	return out, nil
}

func containsStr(arr []string, val string) bool {
	for _, s := range arr {
		if s == val {
			return true
		}
	}
	return false
}

func handleAPISearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	term := r.URL.Query().Get("term")
	if strings.TrimSpace(term) == "" {
		jsonError(w, http.StatusBadRequest, "search term is required")
		return
	}

	country := r.URL.Query().Get("country")
	accountID := r.URL.Query().Get("accountId")
	if country == "" {
		if acc, ok := accountsManager.GetAccount(accountID); ok && acc.StoreFront != "" {
			if cc, err := appstore.CountryCodeFromStoreFront(acc.StoreFront); err == nil {
				country = cc
			}
		}
	}
	if country == "" {
		country = "US"
	}

	limit := 25
	if l := r.URL.Query().Get("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil && n > 0 && n <= 50 {
			limit = n
		}
	}

	results, err := searchAppStore(term, country, limit)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"count":   len(results),
		"results": results,
		"country": country,
	})
}

// handleAPICheckLicense checks whether the active or specified Apple ID owns a license
// for the given app, without downloading any IPA bytes.
func handleAPICheckLicense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	appIDStr := strings.TrimSpace(r.URL.Query().Get("id"))
	bundleID := strings.TrimSpace(r.URL.Query().Get("bundleId"))
	accountID := strings.TrimSpace(r.URL.Query().Get("accountId"))

	if appIDStr == "" && bundleID == "" {
		jsonError(w, http.StatusBadRequest, "id or bundleId is required")
		return
	}

	// 1. Resolve Account
	var accRec *AccountRecord
	if accountID != "" {
		var ok bool
		accRec, ok = accountsManager.GetAccount(accountID)
		if !ok || accRec == nil {
			jsonError(w, http.StatusBadRequest, "Tài khoản Apple ID không tồn tại")
			return
		}
	} else {
		var ok bool
		accRec, ok = accountsManager.ActiveAccount()
		if !ok || accRec == nil {
			jsonResponse(w, http.StatusOK, map[string]any{
				"success":       true,
				"authenticated": false,
				"owned":         false,
				"message":       "Chưa đăng nhập tài khoản Apple ID",
			})
			return
		}
	}

	appAcc := accRec.AsAppStoreAccount()
	client, err := accountsManager.ClientForAccount(accRec.ID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Khởi tạo App Store client thất bại: "+err.Error())
		return
	}

	// 2. Resolve App Metadata
	var target decryptTarget
	if appIDStr != "" {
		target.appId = appIDStr
	} else {
		target.bundleId = bundleID
	}

	app, err := lookupTargetApp(client, appAcc, target)
	if err != nil {
		jsonError(w, http.StatusNotFound, "Không tìm thấy ứng dụng trên App Store: "+err.Error())
		return
	}

	isFree := app.Price == 0

	// 3. Test License / Ownership via PrepareDownload (queries license authorization without downloading)
	_, err = client.PrepareDownload(appAcc, app, "")
	if err == nil {
		jsonResponse(w, http.StatusOK, map[string]any{
			"success":       true,
			"authenticated": true,
			"owned":         true,
			"isFree":        isFree,
			"price":         app.Price,
			"status":        "owned",
			"appId":         app.ID,
			"bundleId":      app.BundleID,
			"message":       "Tài khoản của bạn đã sở hữu ứng dụng này",
		})
		return
	}

	if errors.Is(err, appstore.ErrLicenseRequired) || strings.Contains(err.Error(), "9610") || strings.Contains(err.Error(), "license required") {
		if isFree {
			jsonResponse(w, http.StatusOK, map[string]any{
				"success":        true,
				"authenticated":  true,
				"owned":          false,
				"isFree":         true,
				"price":          0,
				"canAutoAcquire": true,
				"status":         "free_unclaimed",
				"appId":          app.ID,
				"bundleId":       app.BundleID,
				"message":        "Ứng dụng miễn phí chưa nhận. ipadecrypt sẽ tự động liên kết bản quyền khi tải.",
			})
		} else {
			jsonResponse(w, http.StatusOK, map[string]any{
				"success":        true,
				"authenticated":  true,
				"owned":          false,
				"isFree":         false,
				"price":          app.Price,
				"canAutoAcquire": false,
				"status":         "not_owned",
				"appId":          app.ID,
				"bundleId":       app.BundleID,
				"message":        "Tài khoản chưa mua ứng dụng trả phí này trên App Store.",
			})
		}
		return
	}

	if errors.Is(err, appstore.ErrModernDownloadUnsupported) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"success":       true,
			"authenticated": true,
			"owned":         true,
			"isFree":        isFree,
			"price":         app.Price,
			"status":        "owned_attestation",
			"appId":         app.ID,
			"bundleId":      app.BundleID,
			"message":       "Tài khoản đã sở hữu bản quyền ứng dụng này.",
		})
		return
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":       true,
		"authenticated": true,
		"owned":         false,
		"isFree":        isFree,
		"price":         app.Price,
		"status":        "unknown",
		"appId":         app.ID,
		"bundleId":      app.BundleID,
		"message":       err.Error(),
	})
}
