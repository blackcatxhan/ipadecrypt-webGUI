@echo off
title ipadecrypt GUI — FairPlay App Store Decrypter for Windows
cd /d "%~dp0"

echo =======================================================
echo          ipadecrypt GUI for Windows
echo   End-to-End FairPlay IPA Decrypter Web Interface
echo =======================================================
echo.

:: Check if port 54322 is already in use
netstat -ano | findstr :54322 | findstr LISTENING >nul 2>nul
if %ERRORLEVEL% equ 0 (
    echo [ERROR] Port 54322 is already in use!
    echo.
    echo Another instance of ipadecrypt GUI or another application
    echo is already using port 54322.
    echo.
    echo Please close the other application and try again,
    echo or wait a moment for it to fully shut down.
    echo.
    pause
    exit /b 1
)

echo Checking for ipadecrypt binary...

if exist "ipadecrypt.exe" (
    echo Found ipadecrypt.exe. Starting GUI server...
    ipadecrypt.exe gui --port 54322
    goto :exit
)

if exist "bin\ipadecrypt.exe" (
    echo Found bin\ipadecrypt.exe. Starting GUI server...
    bin\ipadecrypt.exe gui --port 54322
    goto :exit
)

where ipadecrypt >nul 2>nul
if %ERRORLEVEL% equ 0 (
    echo Found ipadecrypt in PATH. Starting GUI server...
    ipadecrypt gui --port 54322
    goto :exit
)

echo [ERROR] ipadecrypt.exe was not found.
echo Please run build.bat first or place ipadecrypt.exe in this folder.
pause
exit /b 1

:exit
echo.
if %ERRORLEVEL% neq 0 (
    echo.
    echo =======================================================
    echo  An error occurred (exit code: %ERRORLEVEL%)
    echo  Read the output above for details.
    echo =======================================================
)
echo.
pause
