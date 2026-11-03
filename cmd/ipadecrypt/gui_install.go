package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type ConnectedDevice struct {
	UDID string `json:"udid"`
	Name string `json:"name,omitempty"`
}

type InstallJob struct {
	ID        string  `json:"id"`
	UDID      string  `json:"udid"`
	FileName  string  `json:"fileName"`
	Status    string  `json:"status"` // "installing", "completed", "error"
	Progress  float64 `json:"progress"`
	Message   string  `json:"message"`
	Error     string  `json:"error,omitempty"`
	CreatedAt int64   `json:"createdAt"`
}

var (
	guiInstallJobsMu sync.RWMutex
	guiInstallJobs   = make(map[string]*InstallJob)
)

func findDeviceTool(name string) string {
	candidates := []string{
		filepath.Join("tools", name+".exe"),
		filepath.Join("tools", name),
		name + ".exe",
		name,
	}

	for _, cand := range candidates {
		if path, err := exec.LookPath(cand); err == nil {
			return path
		}
	}
	return ""
}

func listConnectedDevices() ([]ConnectedDevice, error) {
	ideviceID := findDeviceTool("idevice_id")
	if ideviceID == "" {
		return nil, fmt.Errorf("không tìm thấy idevice_id (cần libimobiledevice trong thư mục tools)")
	}

	cmd := exec.Command(ideviceID, "-l")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("chạy idevice_id thất bại: %w", err)
	}

	var devices []ConnectedDevice
	scanner := bufio.NewScanner(&stdout)
	for scanner.Scan() {
		udid := strings.TrimSpace(scanner.Text())
		if udid != "" {
			name := getDeviceName(udid)
			devices = append(devices, ConnectedDevice{
				UDID: udid,
				Name: name,
			})
		}
	}

	return devices, nil
}

func getDeviceName(udid string) string {
	idevicename := findDeviceTool("idevicename")
	if idevicename != "" {
		cmd := exec.Command(idevicename, "-u", udid)
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err == nil {
			if s := strings.TrimSpace(out.String()); s != "" {
				return s
			}
		}
	}
	return udid
}

func handleAPIInstallDevices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	devices, err := listConnectedDevices()
	toolsAvailable := findDeviceTool("ideviceinstaller") != ""

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":        true,
		"toolsAvailable": toolsAvailable,
		"devices":        devices,
		"error": func() string {
			if err != nil {
				return err.Error()
			}
			return ""
		}(),
	})
}

func handleAPIInstallApp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		UDID     string `json:"udid"`
		FileName string `json:"fileName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.UDID == "" || req.FileName == "" {
		jsonError(w, http.StatusBadRequest, "udid and fileName are required")
		return
	}

	installer := findDeviceTool("ideviceinstaller")
	if installer == "" {
		jsonError(w, http.StatusInternalServerError, "Không tìm thấy công cụ ideviceinstaller trong thư mục tools")
		return
	}

	filePath := filepath.Join(getOutputDir(), req.FileName)

	jobID := newJobID()
	job := &InstallJob{
		ID:        jobID,
		UDID:      req.UDID,
		FileName:  req.FileName,
		Status:    "installing",
		Message:   "Đang cài đặt IPA vào thiết bị...",
		CreatedAt: time.Now().Unix(),
	}

	guiInstallJobsMu.Lock()
	guiInstallJobs[jobID] = job
	guiInstallJobsMu.Unlock()

	go func() {
		cmd := exec.Command(installer, "-u", req.UDID, "-i", filePath)
		var out bytes.Buffer
		var errOut bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &errOut

		err := cmd.Run()
		guiInstallJobsMu.Lock()
		defer guiInstallJobsMu.Unlock()

		if err != nil {
			job.Status = "error"
			job.Error = strings.TrimSpace(errOut.String())
			if job.Error == "" {
				job.Error = strings.TrimSpace(out.String())
			}
			if job.Error == "" {
				job.Error = err.Error()
			}
			job.Message = "Cài đặt thất bại"
		} else {
			job.Status = "completed"
			job.Progress = 100
			job.Message = "Đã cài đặt thành công lên thiết bị!"
		}
	}()

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"jobId":   jobID,
		"job":     job,
	})
}

func handleAPIInstallStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		jsonError(w, http.StatusBadRequest, "id is required")
		return
	}

	guiInstallJobsMu.RLock()
	job, ok := guiInstallJobs[id]
	guiInstallJobsMu.RUnlock()

	if !ok || job == nil {
		jsonError(w, http.StatusNotFound, "job not found")
		return
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"job":     job,
	})
}
