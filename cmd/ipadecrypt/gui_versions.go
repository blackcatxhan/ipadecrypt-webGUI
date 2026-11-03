package main

import (
	"net/http"
	"strings"
	"sync"
)

type VersionItem struct {
	VersionID        string `json:"versionId"`
	DisplayVersion   string `json:"displayVersion,omitempty"`
	BundleVersion    string `json:"bundleVersion,omitempty"`
	ReleaseDate      string `json:"releaseDate,omitempty"`
	ReleaseTimestamp int64  `json:"releaseTimestamp,omitempty"`
	FileSize         int64  `json:"fileSize,omitempty"`
	Current          bool   `json:"current"`
}

type cachedVersionMeta struct {
	DisplayVersion string
	BundleVersion  string
	ReleaseDate    string
	ReleaseTime    int64
	FileSize       int64
}

var (
	guiVersionCacheMu sync.RWMutex
	guiVersionCache   = make(map[string]map[string]cachedVersionMeta) // appTrackId -> versionId -> meta
)

func handleAPIVersions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	raw := r.URL.Query().Get("target")
	if strings.TrimSpace(raw) == "" {
		jsonError(w, http.StatusBadRequest, "target (bundle ID or App Store ID) is required")
		return
	}

	accountID := r.URL.Query().Get("accountId")
	accRec, ok := accountsManager.GetAccount(accountID)
	if !ok || accRec == nil || accRec.PasswordToken == "" {
		jsonError(w, http.StatusUnauthorized, "Vui lòng đăng nhập Apple ID trước khi xem lịch sử phiên bản")
		return
	}

	client, err := accountsManager.ClientForAccount(accRec.ID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	appAcc := accRec.AsAppStoreAccount()

	target, err := parseDecryptArg(raw)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	app, err := lookupTargetApp(client, appAcc, target)
	if err != nil {
		jsonError(w, http.StatusNotFound, "Không tìm thấy ứng dụng trên App Store: "+err.Error())
		return
	}

	// List versions
	listOut, err := client.ListVersions(appAcc, app)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Lỗi lấy danh sách phiên bản: "+err.Error())
		return
	}

	appTrackKey := app.BundleID
	guiVersionCacheMu.RLock()
	cachedAppMeta := guiVersionCache[appTrackKey]
	guiVersionCacheMu.RUnlock()

	// ExternalVersionIDs is ordered oldest -> newest by Apple
	// Reverse so newest is first in UI
	ids := make([]string, len(listOut.ExternalVersionIDs))
	for i, id := range listOut.ExternalVersionIDs {
		ids[len(listOut.ExternalVersionIDs)-1-i] = id
	}

	items := make([]VersionItem, 0, len(ids))
	for _, id := range ids {
		item := VersionItem{
			VersionID: id,
			Current:   id == listOut.LatestExternalVersionID,
		}
		if cachedAppMeta != nil {
			if meta, found := cachedAppMeta[id]; found {
				item.DisplayVersion = meta.DisplayVersion
				item.BundleVersion = meta.BundleVersion
				item.ReleaseDate = meta.ReleaseDate
				item.ReleaseTimestamp = meta.ReleaseTime
				item.FileSize = meta.FileSize
			}
		}
		items = append(items, item)
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":       true,
		"appName":       app.Name,
		"bundleId":      app.BundleID,
		"appId":         app.ID,
		"latestVersion": app.Version,
		"latestExtId":   listOut.LatestExternalVersionID,
		"count":         len(items),
		"versions":      items,
	})
}

func handleAPIVersionMetadata(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	targetRaw := r.URL.Query().Get("target")
	versionID := r.URL.Query().Get("versionId")
	if targetRaw == "" || versionID == "" {
		jsonError(w, http.StatusBadRequest, "target and versionId are required")
		return
	}

	accountID := r.URL.Query().Get("accountId")
	accRec, ok := accountsManager.GetAccount(accountID)
	if !ok || accRec == nil || accRec.PasswordToken == "" {
		jsonError(w, http.StatusUnauthorized, "Vui lòng đăng nhập Apple ID")
		return
	}

	target, err := parseDecryptArg(targetRaw)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	client, err := accountsManager.ClientForAccount(accRec.ID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	appAcc := accRec.AsAppStoreAccount()
	app, err := lookupTargetApp(client, appAcc, target)
	if err != nil {
		jsonError(w, http.StatusNotFound, err.Error())
		return
	}

	// Check cache
	appKey := app.BundleID
	guiVersionCacheMu.RLock()
	if appM, ok := guiVersionCache[appKey]; ok {
		if meta, found := appM[versionID]; found {
			guiVersionCacheMu.RUnlock()
			jsonResponse(w, http.StatusOK, map[string]any{
				"success":        true,
				"versionId":      versionID,
				"displayVersion": meta.DisplayVersion,
				"bundleVersion":  meta.BundleVersion,
				"releaseDate":    meta.ReleaseDate,
				"fileSize":       meta.FileSize,
			})
			return
		}
	}
	guiVersionCacheMu.RUnlock()

	// Fetch from Apple
	meta, err := client.GetVersionMetadata(appAcc, app, versionID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Lỗi lấy thông tin phiên bản: "+err.Error())
		return
	}

	dateStr := ""
	var dateTs int64
	if !meta.ReleaseDate.IsZero() {
		dateStr = meta.ReleaseDate.Format("2006-01-02 15:04:05")
		dateTs = meta.ReleaseDate.Unix()
	}

	// Save to cache
	guiVersionCacheMu.Lock()
	if guiVersionCache[appKey] == nil {
		guiVersionCache[appKey] = make(map[string]cachedVersionMeta)
	}
	guiVersionCache[appKey][versionID] = cachedVersionMeta{
		DisplayVersion: meta.DisplayVersion,
		BundleVersion:  meta.BundleVersion,
		ReleaseDate:    dateStr,
		ReleaseTime:    dateTs,
		FileSize:       meta.FileSize,
	}
	guiVersionCacheMu.Unlock()

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":        true,
		"versionId":      versionID,
		"displayVersion": meta.DisplayVersion,
		"bundleVersion":  meta.BundleVersion,
		"releaseDate":    dateStr,
		"fileSize":       meta.FileSize,
	})
}
