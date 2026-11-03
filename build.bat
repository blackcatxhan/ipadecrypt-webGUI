@echo off
setlocal enabledelayedexpansion
chcp 65001 >nul
title ipadecrypt GUI - Windows Builder
cd /d "%~dp0"

echo ======================================================================
echo           ipadecrypt WebGUI - Windows Executable Builder
echo ======================================================================
echo.

:: 1. Check prerequisites
echo [1/3] Checking required build tools...

where go >nul 2>nul
if %ERRORLEVEL% neq 0 (
    echo [ERROR] Go compiler is not installed or not found in PATH!
    echo Please install Go for Windows from: https://go.dev/dl/
    pause
    exit /b 1
)

for /f "tokens=*" %%v in ('go version') do set "GO_VER=%%v"
echo  - %GO_VER%
echo.

:: 2. Build
echo [2/3] Compiling ipadecrypt.exe with WebGUI embedded...
go build -ldflags "-s -w" -o ipadecrypt.exe ./cmd/ipadecrypt
if %ERRORLEVEL% neq 0 (
    echo.
    echo [ERROR] Build failed! Check the errors above.
    pause
    exit /b 1
)

echo.
echo [3/3] Build completed successfully!
echo Executable generated: ipadecrypt.exe
echo.
echo You can now run:
echo   - ipadecrypt-gui.bat (to launch the WebGUI in your browser)
echo   - .\ipadecrypt.exe gui
echo.
pause
