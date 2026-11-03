# ipadecrypt GUI - PowerShell Launcher
$Host.UI.RawUI.WindowTitle = "ipadecrypt GUI — FairPlay App Store Decrypter"
Set-Location $PSScriptRoot

Write-Host "=======================================================" -ForegroundColor Cyan
Write-Host "         ipadecrypt GUI for Windows" -ForegroundColor Cyan
Write-Host "  End-to-End FairPlay IPA Decrypter Web Interface" -ForegroundColor DarkCyan
Write-Host "=======================================================" -ForegroundColor Cyan
Write-Host ""

$port = 54322

# Check if port is in use
$portInUse = Get-NetTCPConnection -LocalPort $port -State Listen -ErrorAction SilentlyContinue
if ($portInUse) {
    Write-Host "[ERROR] Port $port is already in use!" -ForegroundColor Red
    Write-Host "Please close the other application or process using port $port."
    Read-Host "Press Enter to exit..."
    exit 1
}

$exe = $null
if (Test-Path ".\ipadecrypt.exe") {
    $exe = ".\ipadecrypt.exe"
} elseif (Test-Path ".\bin\ipadecrypt.exe") {
    $exe = ".\bin\ipadecrypt.exe"
} elseif (Get-Command "ipadecrypt" -ErrorAction SilentlyContinue) {
    $exe = "ipadecrypt"
}

if (-not $exe) {
    Write-Host "[ERROR] ipadecrypt.exe was not found." -ForegroundColor Red
    Write-Host "Please run .\build.bat to build the executable."
    Read-Host "Press Enter to exit..."
    exit 1
}

Write-Host "Starting ipadecrypt GUI server on port $port..." -ForegroundColor Green
& $exe gui --port $port
