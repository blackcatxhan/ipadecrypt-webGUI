package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/londek/ipadecrypt/internal/credmint"
	"github.com/londek/ipadecrypt/internal/device"
)

type DeviceConfigView struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	USB      bool   `json:"usb"`
	UDID     string `json:"udid"`
	AuthKind string `json:"authKind"` // "password" or "key"
	Password string `json:"password,omitempty"`
	KeyPath  string `json:"keyPath,omitempty"`
}

type DeviceTestCheck struct {
	Name   string `json:"name"`
	Passed bool   `json:"passed"`
	Detail string `json:"detail"`
}

type DeviceTestResult struct {
	Connected bool              `json:"connected"`
	Ready     bool              `json:"ready"`
	Checks    []DeviceTestCheck `json:"checks"`
	Error     string            `json:"error,omitempty"`
}

func handleAPIDeviceConfig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		d := globalConfig.Device
		port := d.Port
		if port == 0 {
			port = 22
		}
		user := d.User
		if user == "" {
			user = "root"
		}
		authKind := d.Auth.Kind
		if authKind == "" {
			authKind = "password"
		}

		jsonResponse(w, http.StatusOK, map[string]any{
			"success": true,
			"config": DeviceConfigView{
				Host:     d.Host,
				Port:     port,
				User:     user,
				USB:      d.USB,
				UDID:     d.UDID,
				AuthKind: authKind,
				Password: d.Auth.Password,
				KeyPath:  d.Auth.KeyPath,
			},
		})

	case http.MethodPost:
		var req DeviceConfigView
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			jsonError(w, http.StatusBadRequest, "invalid request body")
			return
		}

		globalConfig.Device.Host = strings.TrimSpace(req.Host)
		if req.Port <= 0 {
			req.Port = 22
		}
		globalConfig.Device.Port = req.Port
		if req.User == "" {
			req.User = "root"
		}
		globalConfig.Device.User = strings.TrimSpace(req.User)
		globalConfig.Device.USB = req.USB
		globalConfig.Device.UDID = strings.TrimSpace(req.UDID)

		if req.AuthKind == "" {
			req.AuthKind = "password"
		}
		globalConfig.Device.Auth.Kind = req.AuthKind
		if req.Password != "" {
			globalConfig.Device.Auth.Password = req.Password
		}
		globalConfig.Device.Auth.KeyPath = req.KeyPath

		if err := globalConfig.Save(); err != nil {
			jsonError(w, http.StatusInternalServerError, "Lỗi lưu cấu hình: "+err.Error())
			return
		}

		jsonResponse(w, http.StatusOK, map[string]any{
			"success": true,
			"message": "Đã lưu cấu hình thiết bị thành công",
		})

	default:
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func handleAPIDeviceTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	devCfg := globalConfig.Device
	if devCfg.Host == "" && !devCfg.USB {
		jsonResponse(w, http.StatusOK, map[string]any{
			"success": false,
			"error":   "Chưa nhập địa chỉ IP hoặc chọn chế độ USB cho thiết bị",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dev, err := device.Connect(ctx, devCfg)
	if err != nil {
		jsonResponse(w, http.StatusOK, map[string]any{
			"success": false,
			"error":   fmt.Sprintf("Không thể kết nối đến thiết bị (SSH): %v", err),
		})
		return
	}
	defer dev.Close()

	checks := make([]DeviceTestCheck, 0)
	allPassed := true

	// Check 1: OpenSSH connection
	checks = append(checks, DeviceTestCheck{
		Name:   "Root OpenSSH",
		Passed: true,
		Detail: "Kết nối SSH thành công với quyền root",
	})

	// Check 2: AppSync Unified
	appSyncPath, err := dev.LocateAppSync()
	if err != nil || appSyncPath == "" {
		checks = append(checks, DeviceTestCheck{
			Name:   "AppSync Unified",
			Passed: false,
			Detail: "Chưa cài đặt (cần cài đặt từ Sileo: https://lukezgd.github.io/repo)",
		})
		allPassed = false
	} else {
		checks = append(checks, DeviceTestCheck{
			Name:   "AppSync Unified",
			Passed: true,
			Detail: "Đã tìm thấy tại " + appSyncPath,
		})
	}

	// Check 3: appinst
	appinstPath, err := dev.LocateAppinst()
	if err != nil || appinstPath == "" {
		checks = append(checks, DeviceTestCheck{
			Name:   "appinst",
			Passed: false,
			Detail: "Chưa cài đặt (cần cài đặt từ Sileo)",
		})
		allPassed = false
	} else {
		checks = append(checks, DeviceTestCheck{
			Name:   "appinst",
			Passed: true,
			Detail: "Đã tìm thấy tại " + appinstPath,
		})
	}

	// Check 4: Upload helper & verify
	helperPath, err := dev.EnsureHelper()
	if err != nil {
		checks = append(checks, DeviceTestCheck{
			Name:   "Decrypt Helper",
			Passed: false,
			Detail: "Lỗi tải helper lên thiết bị: " + err.Error(),
		})
		allPassed = false
	} else if err := dev.VerifyHelper(helperPath); err != nil {
		checks = append(checks, DeviceTestCheck{
			Name:   "Decrypt Helper",
			Passed: false,
			Detail: "Lỗi thực thi helper: " + err.Error(),
		})
		allPassed = false
	} else {
		checks = append(checks, DeviceTestCheck{
			Name:   "Decrypt Helper",
			Passed: true,
			Detail: "Helper sẵn sàng tại " + helperPath,
		})
	}

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":   true,
		"connected": true,
		"ready":     allPassed,
		"checks":    checks,
	})
}

func handleAPIDeviceMint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		AccountID string `json:"accountId"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	accRec, ok := accountsManager.GetAccount(req.AccountID)
	if !ok || accRec == nil {
		jsonError(w, http.StatusBadRequest, "Không tìm thấy tài khoản Apple ID")
		return
	}

	if accRec.DSID == "" {
		jsonError(w, http.StatusBadRequest, "Tài khoản chưa có mã nhận dạng (DSID). Vui lòng đăng nhập lại.")
		return
	}

	devCfg := globalConfig.Device
	if devCfg.Host == "" && !devCfg.USB {
		jsonError(w, http.StatusBadRequest, "Chưa cấu hình thiết bị Jailbreak. Vui lòng vào tab Thiết bị để nhập thông tin kết nối.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	credential, err := credmint.Mint(ctx, devCfg, accRec.DSID)
	if err != nil {
		if errors.Is(err, credmint.ErrBinariesUnavailable) {
			jsonError(w, http.StatusInternalServerError, "Bản build này không chứa các file nhị phân của thiết bị để tạo chứng chỉ.")
			return
		}
		jsonError(w, http.StatusInternalServerError, "Lỗi tạo chứng chỉ FairPlay: "+err.Error())
		return
	}

	if err := saveCredential(globalPaths, accRec.DSID, credential); err != nil {
		jsonError(w, http.StatusInternalServerError, "Lỗi lưu chứng chỉ: "+err.Error())
		return
	}

	credPath, _ := globalPaths.CredentialPath(accRec.DSID)

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":        true,
		"message":        fmt.Sprintf("Đã tạo và lưu thành công chứng chỉ FairPlay 16-byte cho tài khoản %s!", accRec.Email),
		"dsid":           accRec.DSID,
		"credentialFile": credPath,
	})
}
