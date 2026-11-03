package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type BatchItem struct {
	ID        string `json:"id"`
	Target    string `json:"target"`
	Status    string `json:"status"` // "pending", "running", "completed", "error", "skipped"
	JobID     string `json:"jobId,omitempty"`
	AppName   string `json:"appName,omitempty"`
	BundleID  string `json:"bundleId,omitempty"`
	Error     string `json:"error,omitempty"`
	OutputIPA string `json:"outputIpa,omitempty"`
}

type BatchTask struct {
	ID        string       `json:"id"`
	AccountID string       `json:"accountId"`
	Total     int          `json:"total"`
	Completed int          `json:"completed"`
	Failed    int          `json:"failed"`
	Skipped   int          `json:"skipped"`
	Status    string       `json:"status"` // "running", "paused", "completed", "stopped"
	Items     []*BatchItem `json:"items"`
	CreatedAt int64        `json:"createdAt"`

	mu     sync.RWMutex       `json:"-"`
	ctx    context.Context    `json:"-"`
	cancel context.CancelFunc `json:"-"`
	pause  chan struct{}      `json:"-"`
	skip   chan struct{}      `json:"-"`
}

var (
	guiBatchTrackerMu sync.RWMutex
	guiActiveBatch    *BatchTask
)

func parseBatchInput(text string) []string {
	scanner := bufio.NewScanner(strings.NewReader(text))
	var targets []string
	seen := make(map[string]bool)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		// Support "Name: AppID" format or bare bundle ID / App ID
		if parts := strings.Split(line, ":"); len(parts) >= 2 {
			candidate := strings.TrimSpace(parts[len(parts)-1])
			if candidate != "" {
				line = candidate
			}
		}

		if !seen[line] {
			seen[line] = true
			targets = append(targets, line)
		}
	}

	return targets
}

func handleAPIBatchStart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		Text      string `json:"text"`
		AccountID string `json:"accountId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	targets := parseBatchInput(req.Text)
	if len(targets) == 0 {
		jsonError(w, http.StatusBadRequest, "Không tìm thấy danh sách mục tiêu hợp lệ")
		return
	}

	accRec, ok := accountsManager.GetAccount(req.AccountID)
	if !ok || accRec == nil {
		jsonError(w, http.StatusBadRequest, "Chưa chọn tài khoản Apple ID")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	batch := &BatchTask{
		ID:        newJobID(),
		AccountID: accRec.ID,
		Total:     len(targets),
		Status:    "running",
		Items:     make([]*BatchItem, 0, len(targets)),
		CreatedAt: time.Now().Unix(),
		ctx:       ctx,
		cancel:    cancel,
		skip:      make(chan struct{}, 1),
	}

	for i, t := range targets {
		batch.Items = append(batch.Items, &BatchItem{
			ID:     fmt.Sprintf("%d", i+1),
			Target: t,
			Status: "pending",
		})
	}

	guiBatchTrackerMu.Lock()
	if guiActiveBatch != nil && guiActiveBatch.Status == "running" {
		guiActiveBatch.cancel()
	}
	guiActiveBatch = batch
	guiBatchTrackerMu.Unlock()

	go runBatchLoop(batch)

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"batchId": batch.ID,
		"batch":   batch,
	})
}

func runBatchLoop(batch *BatchTask) {
	for _, item := range batch.Items {
		select {
		case <-batch.ctx.Done():
			batch.Status = "stopped"
			return
		default:
		}

		item.Status = "running"

		jobCtx, jobCancel := context.WithCancel(batch.ctx)
		job := &DecryptJob{
			ID:           newJobID(),
			AccountID:    batch.AccountID,
			BundleID:     item.Target,
			Status:       "queued",
			Stage:        "Đang xử lý mục tiêu hàng loạt...",
			CreatedAt:    time.Now().Unix(),
			ctx:          jobCtx,
			cancel:       jobCancel,
		}
		item.JobID = job.ID
		guiJobs.add(job)

		runDecryptJob(job)

		if job.Status == "completed" {
			item.Status = "completed"
			item.AppName = job.AppName
			item.BundleID = job.BundleID
			item.OutputIPA = job.OutputName
			batch.mu.Lock()
			batch.Completed++
			batch.mu.Unlock()
		} else {
			item.Status = "error"
			item.Error = job.Error
			batch.mu.Lock()
			batch.Failed++
			batch.mu.Unlock()
		}

		time.Sleep(1 * time.Second)
	}

	batch.Status = "completed"
}

func handleAPIBatchStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	guiBatchTrackerMu.RLock()
	batch := guiActiveBatch
	guiBatchTrackerMu.RUnlock()

	if batch == nil {
		jsonResponse(w, http.StatusOK, map[string]any{
			"success": true,
			"batch":   nil,
		})
		return
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"batch":   batch,
	})
}

func handleAPIBatchStop(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	guiBatchTrackerMu.Lock()
	if guiActiveBatch != nil {
		guiActiveBatch.cancel()
		guiActiveBatch.Status = "stopped"
	}
	guiBatchTrackerMu.Unlock()

	jsonResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Đã dừng tiến trình giải mã hàng loạt",
	})
}
