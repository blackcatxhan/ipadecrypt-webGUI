package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/londek/ipadecrypt/internal/config"
	"github.com/londek/ipadecrypt/internal/tui"
	"github.com/spf13/cobra"
)

//go:embed gui_assets/*
var embeddedAssets embed.FS

var (
	guiPort      int    = 54322
	guiHost      string = "127.0.0.1"
	guiNoOpen    bool   = false
	guiOutputDir string = "./decrypted"

	globalConfig *config.Config
	globalPaths  *config.Paths
)

func jsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func jsonError(w http.ResponseWriter, status int, message string) {
	jsonResponse(w, status, map[string]any{
		"success": false,
		"error":   message,
	})
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	_ = cmd.Start()
}

func handleAPIStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	activeAcc, hasActive := accountsManager.ActiveAccount()
	views, activeID := accountsManager.Views()

	jsonResponse(w, http.StatusOK, map[string]any{
		"success":   true,
		"version":   resolveVersion(),
		"activeId":  activeID,
		"hasActive": hasActive,
		"activeAccount": func() any {
			if !hasActive || activeAcc == nil {
				return nil
			}
			for _, v := range views {
				if v.ID == activeID {
					return v
				}
			}
			return nil
		}(),
		"accountsCount": len(views),
		"deviceConfigured": globalConfig.Device.Host != "" || globalConfig.Device.USB,
		"outputDir":        getOutputDir(),
	})
}

func guiHandler(cmd *cobra.Command, args []string) {
	// Prepend tools folder to PATH so iproxy and ideviceinstaller work automatically
	setupToolsPath()

	cfg, paths, err := loadConfigOrDefault(rootDirOverride)
	if err != nil {
		tui.Err("Lỗi tải cấu hình: %v", err)
		return
	}

	globalConfig = cfg
	globalPaths = paths

	// Ensure output directory exists
	_ = os.MkdirAll(getOutputDir(), 0o755)

	// Initialize Multi-Account manager
	if _, err := initAccountManager(cfg, paths); err != nil {
		tui.Err("Lỗi khởi tạo quản lý tài khoản: %v", err)
		return
	}

	// Setup HTTP Router
	mux := http.NewServeMux()

	// Static Assets
	assetSubFS, err := fs.Sub(embeddedAssets, "gui_assets")
	if err != nil {
		tui.Err("Lỗi nạp giao diện tĩnh: %v", err)
		return
	}
	fileServer := http.FileServer(http.FS(assetSubFS))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If path is root or asset
		p := r.URL.Path
		if p == "/" || p == "/index.html" {
			fileServer.ServeHTTP(w, r)
			return
		}

		// Try serving file directly
		f, err := assetSubFS.Open(strings.TrimPrefix(p, "/"))
		if err == nil {
			_ = f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}

		// Fallback to index.html for SPA routes
		r.URL.Path = "/"
		fileServer.ServeHTTP(w, r)
	})

	// API Routes
	mux.HandleFunc("/api/status", handleAPIStatus)

	// Accounts
	mux.HandleFunc("/api/accounts", handleAPIAccounts)
	mux.HandleFunc("/api/accounts/activate", handleAPIActivateAccount)
	mux.HandleFunc("/api/accounts/remove", handleAPIRemoveAccount)
	mux.HandleFunc("/api/auth/login", handleAPILogin)

	// Search & Versions
	mux.HandleFunc("/api/search", handleAPISearch)
	mux.HandleFunc("/api/app/check-license", handleAPICheckLicense)
	mux.HandleFunc("/api/versions", handleAPIVersions)
	mux.HandleFunc("/api/version-metadata", handleAPIVersionMetadata)

	// Decryption & Jobs
	mux.HandleFunc("/api/decrypt", handleAPIDecrypt)
	mux.HandleFunc("/api/decrypt/status", handleAPIDecryptStatus)
	mux.HandleFunc("/api/decrypt/active", handleAPIDecryptActive)
	mux.HandleFunc("/api/decrypt/cancel", handleAPIDecryptCancel)
	mux.HandleFunc("/api/decrypt/local", handleAPIDecryptLocal)

	// Batch
	mux.HandleFunc("/api/batch/start", handleAPIBatchStart)
	mux.HandleFunc("/api/batch/status", handleAPIBatchStatus)
	mux.HandleFunc("/api/batch/stop", handleAPIBatchStop)

	// Device & Credential Minting
	mux.HandleFunc("/api/device/config", handleAPIDeviceConfig)
	mux.HandleFunc("/api/device/test", handleAPIDeviceTest)
	mux.HandleFunc("/api/device/mint", handleAPIDeviceMint)

	// Output Files
	mux.HandleFunc("/api/files", handleAPIFiles)
	mux.HandleFunc("/api/files/download", handleAPIFilesDownload)
	mux.HandleFunc("/api/files/open", handleAPIFilesOpen)
	mux.HandleFunc("/api/files/delete", handleAPIFilesDelete)

	// Device Installation
	mux.HandleFunc("/api/install/devices", handleAPIInstallDevices)
	mux.HandleFunc("/api/install/app", handleAPIInstallApp)
	mux.HandleFunc("/api/install/status", handleAPIInstallStatus)

	// Find available port
	listener, actualPort, err := listenOnPort(guiHost, guiPort)
	if err != nil {
		tui.Err("Không thể mở cổng mạng: %v", err)
		return
	}

	serverURL := fmt.Sprintf("http://%s:%d", guiHost, actualPort)

	fmt.Println("=======================================================")
	fmt.Printf("   ipadecrypt-webGUI (v%s)\n", resolveVersion())
	fmt.Println("   End-to-end FairPlay IPA Decrypter Web Interface")
	fmt.Println("=======================================================")
	fmt.Println()
	tui.OK("Server đang chạy tại: %s", serverURL)
	absOut, _ := filepath.Abs(getOutputDir())
	tui.Info("Thư mục lưu file IPA giải mã: %s", absOut)
	fmt.Println()
	tui.Info("Nhấn Ctrl+C để dừng server.")

	if !guiNoOpen {
		go func() {
			time.Sleep(300 * time.Millisecond)
			openBrowser(serverURL)
		}()
	}

	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println()
		tui.Info("Đang tắt server...")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
		_ = listener.Close()
	}()

	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		tui.Err("Lỗi máy chủ: %v", err)
	}
}

func listenOnPort(host string, startPort int) (net.Listener, int, error) {
	for port := startPort; port < startPort+20; port++ {
		addr := net.JoinHostPort(host, strconv.Itoa(port))
		l, err := net.Listen("tcp", addr)
		if err == nil {
			return l, port, nil
		}
	}
	return nil, 0, fmt.Errorf("không tìm thấy cổng khả dụng từ %d đến %d", startPort, startPort+20)
}

func setupToolsPath() {
	var candidates []string

	if wd, err := os.Getwd(); err == nil {
		candidates = append(candidates, filepath.Join(wd, "tools"))
	}
	if exe, err := os.Executable(); err == nil {
		candidates = append(candidates, filepath.Join(filepath.Dir(exe), "tools"))
	}

	pathEnv := os.Getenv("PATH")
	for _, cand := range candidates {
		if info, err := os.Stat(cand); err == nil && info.IsDir() {
			abs, _ := filepath.Abs(cand)
			if !strings.Contains(pathEnv, abs) {
				pathEnv = abs + string(os.PathListSeparator) + pathEnv
			}
		}
	}
	_ = os.Setenv("PATH", pathEnv)
}
