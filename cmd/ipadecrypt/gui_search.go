package main

import (
	"encoding/json"
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

		out = append(out, SearchAppResult{
			ID:               item.TrackID,
			BundleID:         item.BundleID,
			Name:             item.TrackName,
			Version:          item.Version,
			Price:            item.Price,
			FormattedPrice:   item.FormattedPrice,
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
