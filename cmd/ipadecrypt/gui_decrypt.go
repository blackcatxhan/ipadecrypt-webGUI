package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/londek/ipadecrypt/internal/appstore"
	"github.com/londek/ipadecrypt/internal/fairplay"
	"github.com/londek/ipadecrypt/internal/pipeline"
)

type DecryptJob struct {
	ID           string  `json:"id"`
	AccountID    string  `json:"accountId"`
	AccountEmail string  `json:"accountEmail"`
	AppName      string  `json:"appName"`
	BundleID     string  `json:"bundleId"`
	AppID        string  `json:"appId"`
	Version      string  `json:"version"`
	ExtVerID     string  `json:"extVerId,omitempty"`
	Storefront   string  `json:"storefront,omitempty"`
	Status       string  `json:"status"` // "queued", "resolving", "licensing", "downloading", "decrypting", "verifying", "completed", "error", "cancelled"
	Progress     float64 `json:"progress"`
	BytesRead    int64   `json:"bytesRead"`
	TotalBytes   int64   `json:"totalBytes"`
	Speed        string  `json:"speed"`
	Stage        string  `json:"stage"`
	OutputPath   string  `json:"outputPath,omitempty"`
	OutputName   string  `json:"outputName,omitempty"`
	OutputSize   int64   `json:"outputSize,omitempty"`
	Error        string  `json:"error,omitempty"`
	Warning      string  `json:"warning,omitempty"`
	CreatedAt    int64   `json:"createdAt"`
	CompletedAt  int64   `json:"completedAt,omitempty"`

	// Flags
	OnDevice        bool `json:"onDevice"`
	NoOnDevice      bool `json:"noOnDevice"`
	PatchDeviceType bool `json:"patchDeviceType"`
	SkipAppex       bool `json:"skipAppex"`
	NoVerify        bool `json:"noVerify"`
	ExtraVerify     bool `json:"extraVerify"`

	// Local file
	LocalPath string `json:"localPath,omitempty"`

	cancel context.CancelFunc `json:"-"`
	ctx    context.Context    `json:"-"`
}

type jobTracker struct {
	sync.RWMutex
	jobs []*DecryptJob
}

var guiJobs = &jobTracker{
	jobs: make([]*DecryptJob, 0),
}

func (jt *jobTracker) add(j *DecryptJob) {
	jt.Lock()
	defer jt.Unlock()
	jt.jobs = append([]*DecryptJob{j}, jt.jobs...) // newest first
	if len(jt.jobs) > 100 {
		jt.jobs = jt.jobs[:100]
	}
}

func (jt *jobTracker) get(id string) *DecryptJob {
	jt.RLock()
	defer jt.RUnlock()
	for _, j := range jt.jobs {
		if j.ID == id {
			return j
		}
	}
	return nil
}

func (jt *jobTracker) all() []*DecryptJob {
	jt.RLock()
	defer jt.RUnlock()
	res := make([]*DecryptJob, len(jt.jobs))
	copy(res, jt.jobs)
	return res
}

func newJobID() string {
	var b [6]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func (j *DecryptJob) setStatus(status, stage string) {
	j.Status = status
	j.Stage = stage
}

func (j *DecryptJob) setProgress(bytesRead, totalBytes int64, speed string) {
	j.BytesRead = bytesRead
	j.TotalBytes = totalBytes
	j.Speed = speed
	if totalBytes > 0 {
		j.Progress = float64(bytesRead) / float64(totalBytes) * 100.0
		if j.Progress > 100.0 {
			j.Progress = 100.0
		}
	}
}

// runDecryptJob executes a decryption task in the background.
func runDecryptJob(job *DecryptJob) {
	defer func() {
		if r := recover(); r != nil {
			job.Status = "error"
			job.Error = fmt.Sprintf("Panic in decrypt job: %v", r)
			job.CompletedAt = time.Now().Unix()
		}
	}()

	accRec, ok := accountsManager.GetAccount(job.AccountID)
	if !ok || accRec == nil {
		job.Status = "error"
		job.Error = "Tài khoản Apple ID không tồn tại hoặc chưa được chọn"
		job.CompletedAt = time.Now().Unix()
		return
	}

	appAcc := accRec.AsAppStoreAccount()
	client, err := accountsManager.ClientForAccount(accRec.ID)
	if err != nil {
		job.Status = "error"
		job.Error = "Khởi tạo App Store client thất bại: " + err.Error()
		job.CompletedAt = time.Now().Unix()
		return
	}

	if job.Storefront != "" {
		if sf, err := appstore.ResolveStorefront(job.Storefront); err == nil {
			appAcc.StoreFront = sf
		}
	}

	var encPath string
	var appBundleID string
	var appVersion string
	var appName string

	// Step 1: Obtain encrypted IPA
	if job.LocalPath != "" {
		job.setStatus("resolving", "Đang phân tích file IPA cục bộ...")
		bID, ver, err := pipeline.AppInfo(job.LocalPath)
		if err != nil {
			job.Status = "error"
			job.Error = "Đọc file IPA thất bại: " + err.Error()
			job.CompletedAt = time.Now().Unix()
			return
		}
		appBundleID = bID
		appVersion = ver
		appName = filepath.Base(job.LocalPath)
		encPath = job.LocalPath

		job.BundleID = appBundleID
		job.Version = appVersion
		job.AppName = appName
	} else {
		job.setStatus("resolving", "Đang tìm kiếm ứng dụng trên App Store...")
		target, err := parseDecryptArg(job.BundleID)
		if err != nil {
			job.Status = "error"
			job.Error = "Định dạng mục tiêu không hợp lệ: " + err.Error()
			job.CompletedAt = time.Now().Unix()
			return
		}

		app, err := lookupTargetApp(client, appAcc, target)
		if err != nil {
			job.Status = "error"
			job.Error = "Không tìm thấy ứng dụng trên App Store: " + err.Error()
			job.CompletedAt = time.Now().Unix()
			return
		}

		appName = app.Name
		appBundleID = app.BundleID
		appVersion = app.Version

		job.AppName = appName
		job.BundleID = appBundleID
		job.AppID = fmt.Sprintf("%d", app.ID)
		job.Version = appVersion

		job.setStatus("downloading", "Đang chuẩn bị tải về từ App Store...")

		var lastTime time.Time = time.Now()
		var lastBytes int64 = 0

		disposition, err := fetchRemoteEncryptedSource(globalConfig, globalPaths, client, app, job.ExtVerID, func(e authEvent) {
			switch e {
			case authReauth:
				job.setStatus("licensing", "Đang làm mới phiên đăng nhập...")
			case authLicense:
				job.setStatus("licensing", "Đang nhận bản quyền ứng dụng...")
			case authRetryingDownload:
				job.setStatus("downloading", "Đang thử lại tải về...")
			}
		}, func(cur, total int64) {
			now := time.Now()
			dur := now.Sub(lastTime).Seconds()
			speedStr := ""
			if dur >= 0.5 && cur > lastBytes {
				bps := float64(cur-lastBytes) / dur
				speedStr = fmt.Sprintf("%s/s", humanBytes(int64(bps)))
				lastBytes = cur
				lastTime = now
			}
			job.setProgress(cur, total, speedStr)
			job.Stage = fmt.Sprintf("Đang tải IPA: %s / %s", humanBytes(cur), humanBytes(total))
		})

		if err != nil {
			job.Status = "error"
			job.Error = "Tải IPA thất bại: " + err.Error()
			job.CompletedAt = time.Now().Unix()
			return
		}

		encPath = disposition.path
		appVersion = disposition.version
		job.Version = appVersion
	}

	// Step 2: Decrypt
	outDir := getOutputDir()
	sanitizedApp := sanitizeFileName(appName)
	if sanitizedApp == "" {
		sanitizedApp = appBundleID
	}
	outFilename := fmt.Sprintf("%s_%s.decrypted.ipa", sanitizedApp, appVersion)
	outLocal := filepath.Join(outDir, outFilename)

	job.setStatus("decrypting", "Đang giải mã FairPlay...")
	job.Progress = 0

	dsid := accRec.DSID
	var credential []byte
	if !job.OnDevice && dsid != "" {
		cred, err := loadCachedCredential(globalPaths, dsid)
		if err == nil {
			credential = cred
		}
	}

	cleanups := &cleanupStack{}
	defer cleanups.run()

	deviceConfigured := globalConfig.Device.Host != "" || globalConfig.Device.USB

	var offlineSuccess bool = false

	if !job.OnDevice {
		if len(credential) == 16 {
			job.Stage = "Đang giải mã FairPlay offline..."
			outFile, err := os.OpenFile(outLocal, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
			if err != nil {
				job.Status = "error"
				job.Error = "Tạo file đầu ra thất bại: " + err.Error()
				job.CompletedAt = time.Now().Unix()
				return
			}

			cw := &countingWriter{w: outFile, onTick: func(n int64) {
				job.Stage = fmt.Sprintf("Đang ghi IPA giải mã: %s", humanBytes(n))
			}}

			summary, err := fairplay.Decrypt(encPath, cw, credential, job.SkipAppex)
			outFile.Close()

			if err == nil {
				job.setStatus("verifying", "Đang kiểm tra tính toàn vẹn Mach-O...")
				res, vErr := pipeline.Verify(outLocal, encPath, job.SkipAppex)
				if vErr == nil && res.OK() {
					offlineSuccess = true
					job.Stage = fmt.Sprintf("Giải mã thành công %d binaries, %d pages!", summary.Binaries, summary.Pages)
				} else if len(res.StillEncrypted) > 0 && len(res.AllZeroCrypt) == 0 && len(res.Mismatches) == 0 {
					// 32-bit slice needs device fallback
					_ = os.Remove(outLocal)
					if job.NoOnDevice {
						job.Status = "error"
						job.Error = "Ứng dụng chứa slice 32-bit (armv7) chưa được giải mã, nhưng cờ no-on-device được bật"
						job.CompletedAt = time.Now().Unix()
						return
					}
				} else {
					_ = os.Remove(outLocal)
				}
			} else {
				_ = os.Remove(outLocal)
			}
		}
	}

	// Fallback to on-device decryption if offline didn't succeed
	if !offlineSuccess {
		if !deviceConfigured {
			job.Status = "error"
			if len(credential) != 16 {
				job.Error = "Tài khoản chưa có chứng chỉ FairPlay và chưa cấu hình thiết bị Jailbreak. Vui lòng vào tab Thiết bị để Mint Credential."
			} else {
				job.Error = "Giải mã offline không thể xử lý ứng dụng này (thường do chứa mã 32-bit) và chưa có thiết bị Jailbreak được cấu hình để giải mã trực tiếp."
			}
			job.CompletedAt = time.Now().Unix()
			return
		}

		job.setStatus("decrypting", "Đang giải mã trên thiết bị Jailbreak (On-device fallback)...")
		if err := runDeviceDecrypt(globalConfig, cleanups, encPath, appBundleID, appVersion, outLocal); err != nil {
			job.Status = "error"
			job.Error = "Giải mã trên thiết bị thất bại: " + err.Error()
			job.CompletedAt = time.Now().Unix()
			return
		}
	}

	// Output info
	if info, err := os.Stat(outLocal); err == nil {
		job.OutputSize = info.Size()
	}
	job.OutputPath = outLocal
	job.OutputName = filepath.Base(outLocal)
	job.Progress = 100
	job.Status = "completed"
	job.Stage = "Hoàn thành giải mã IPA!"
	job.CompletedAt = time.Now().Unix()
}

func sanitizeFileName(name string) string {
	var b strings.Builder
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' || r == '.' {
			b.WriteRune(r)
		} else if r == ' ' {
			b.WriteRune('_')
		}
	}
	return b.String()
}

// HTTP API Handlers for Decrypt

func handleAPIDecrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		Target          string `json:"target"`
		VersionID       string `json:"versionId"`
		Storefront      string `json:"storefront"`
		OnDevice        bool   `json:"onDevice"`
		NoOnDevice      bool   `json:"noOnDevice"`
		PatchDeviceType bool   `json:"patchDeviceType"`
		SkipAppex       bool   `json:"skipAppex"`
		NoVerify        bool   `json:"noVerify"`
		ExtraVerify     bool   `json:"extraVerify"`
		AccountID       string `json:"accountId"`
		LocalPath       string `json:"localPath"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if strings.TrimSpace(req.Target) == "" && strings.TrimSpace(req.LocalPath) == "" {
		jsonError(w, http.StatusBadRequest, "target or localPath is required")
		return
	}

	accRec, ok := accountsManager.GetAccount(req.AccountID)
	if !ok || accRec == nil {
		jsonError(w, http.StatusBadRequest, "Không tìm thấy tài khoản Apple ID. Vui lòng đăng nhập trước.")
		return
	}

	targetStr := req.Target
	if req.LocalPath != "" {
		targetStr = req.LocalPath
	}

	ctx, cancel := context.WithCancel(context.Background())
	job := &DecryptJob{
		ID:              newJobID(),
		AccountID:       accRec.ID,
		AccountEmail:    accRec.Email,
		BundleID:        targetStr,
		ExtVerID:        req.VersionID,
		Storefront:      req.Storefront,
		OnDevice:        req.OnDevice,
		NoOnDevice:      req.NoOnDevice,
		PatchDeviceType: req.PatchDeviceType,
		SkipAppex:       req.SkipAppex,
		NoVerify:        req.NoVerify,
		ExtraVerify:     req.ExtraVerify,
		LocalPath:       req.LocalPath,
		Status:          "queued",
		Stage:           "Đang chuẩn bị hàng đợi giải mã...",
		CreatedAt:       time.Now().Unix(),
		ctx:             ctx,
		cancel:          cancel,
	}

	guiJobs.add(job)
	go runDecryptJob(job)

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"jobId":   job.ID,
		"job":     job,
	})
}

func handleAPIDecryptStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		jsonError(w, http.StatusBadRequest, "job id is required")
		return
	}

	job := guiJobs.get(id)
	if job == nil {
		jsonError(w, http.StatusNotFound, "job not found")
		return
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"job":     job,
	})
}

func handleAPIDecryptActive(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	jobs := guiJobs.all()
	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"count":   len(jobs),
		"jobs":    jobs,
	})
}

func handleAPIDecryptCancel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ID == "" {
		jsonError(w, http.StatusBadRequest, "job id is required")
		return
	}

	job := guiJobs.get(req.ID)
	if job == nil {
		jsonError(w, http.StatusNotFound, "job not found")
		return
	}

	if job.cancel != nil {
		job.cancel()
	}
	job.Status = "cancelled"
	job.Stage = "Đã hủy bởi người dùng"
	job.CompletedAt = time.Now().Unix()

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Đã hủy tác vụ",
	})
}

func handleAPIDecryptLocal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// 500 MB max in memory
	if err := r.ParseMultipartForm(500 << 20); err != nil {
		jsonError(w, http.StatusBadRequest, "Lỗi nhận file: "+err.Error())
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Không tìm thấy file tải lên")
		return
	}
	defer file.Close()

	if !strings.HasSuffix(strings.ToLower(header.Filename), ".ipa") {
		jsonError(w, http.StatusBadRequest, "File phải có định dạng .ipa")
		return
	}

	cacheDir, _ := globalPaths.CacheDir()
	tempFilePath := filepath.Join(cacheDir, fmt.Sprintf("upload_%d_%s", time.Now().UnixNano(), header.Filename))

	dest, err := os.Create(tempFilePath)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Lỗi tạo file tạm: "+err.Error())
		return
	}
	defer dest.Close()

	if _, err := io.Copy(dest, file); err != nil {
		jsonError(w, http.StatusInternalServerError, "Lỗi ghi file: "+err.Error())
		return
	}

	accountID := r.FormValue("accountId")
	accRec, ok := accountsManager.GetAccount(accountID)
	if !ok || accRec == nil {
		jsonError(w, http.StatusBadRequest, "Chưa chọn tài khoản Apple ID")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	job := &DecryptJob{
		ID:           newJobID(),
		AccountID:    accRec.ID,
		AccountEmail: accRec.Email,
		BundleID:     header.Filename,
		LocalPath:    tempFilePath,
		Status:       "queued",
		Stage:        "Đã nhận file IPA cục bộ, đang chờ giải mã...",
		CreatedAt:    time.Now().Unix(),
		ctx:          ctx,
		cancel:       cancel,
	}

	guiJobs.add(job)
	go runDecryptJob(job)

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"jobId":   job.ID,
		"job":     job,
	})
}
