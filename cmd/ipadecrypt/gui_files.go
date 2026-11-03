package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/londek/ipadecrypt/internal/pipeline"
)

type DecryptedFileInfo struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	Size         int64  `json:"size"`
	SizeHuman    string `json:"sizeHuman"`
	ModifiedAt   int64  `json:"modifiedAt"`
	ModifiedDate string `json:"modifiedDate"`
	BundleID     string `json:"bundleId,omitempty"`
	Version      string `json:"version,omitempty"`
}

func getOutputDir() string {
	if guiOutputDir != "" {
		_ = os.MkdirAll(guiOutputDir, 0o755)
		return guiOutputDir
	}
	dir := filepath.Join(".", "decrypted")
	_ = os.MkdirAll(dir, 0o755)
	return dir
}

func listDecryptedFiles() ([]DecryptedFileInfo, error) {
	outDir := getOutputDir()
	entries, err := os.ReadDir(outDir)
	if err != nil {
		return nil, err
	}

	files := make([]DecryptedFileInfo, 0)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".ipa") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		fullPath := filepath.Join(outDir, entry.Name())
		bundleID, version, _ := pipeline.AppInfo(fullPath)

		files = append(files, DecryptedFileInfo{
			Name:         entry.Name(),
			Path:         fullPath,
			Size:         info.Size(),
			SizeHuman:    humanBytes(info.Size()),
			ModifiedAt:   info.ModTime().Unix(),
			ModifiedDate: info.ModTime().Format("2006-01-02 15:04"),
			BundleID:     bundleID,
			Version:      version,
		})
	}

	// Sort newest first
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModifiedAt > files[j].ModifiedAt
	})

	return files, nil
}

func handleAPIFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	files, err := listDecryptedFiles()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	outDir, _ := filepath.Abs(getOutputDir())

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":   true,
		"outputDir": outDir,
		"count":     len(files),
		"files":     files,
	})
}

func handleAPIFilesDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	filename := r.URL.Query().Get("name")
	if filename == "" || strings.Contains(filename, "..") || strings.ContainsAny(filename, `/\`) {
		jsonError(w, http.StatusBadRequest, "invalid file name")
		return
	}

	filePath := filepath.Join(getOutputDir(), filename)
	info, err := os.Stat(filePath)
	if err != nil {
		jsonError(w, http.StatusNotFound, "file not found")
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"; filename*=UTF-8''%s`,
		filename, url.PathEscape(filename)))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))

	http.ServeContent(w, r, filename, info.ModTime(), f)
}

func handleAPIFilesOpen(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	outDir, err := filepath.Abs(getOutputDir())
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Open folder in Windows Explorer
	cmd := exec.Command("explorer.exe", outDir)
	if err := cmd.Start(); err != nil {
		jsonError(w, http.StatusInternalServerError, "Lỗi mở thư mục: "+err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":   true,
		"outputDir": outDir,
	})
}

func handleAPIFilesDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		jsonError(w, http.StatusBadRequest, "name is required")
		return
	}

	if strings.Contains(req.Name, "..") || strings.ContainsAny(req.Name, `/\`) {
		jsonError(w, http.StatusBadRequest, "invalid file name")
		return
	}

	filePath := filepath.Join(getOutputDir(), req.Name)
	if err := os.Remove(filePath); err != nil {
		jsonError(w, http.StatusInternalServerError, "Lỗi xóa file: "+err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Đã xóa file " + req.Name,
	})
}
