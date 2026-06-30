# install.ps1
$Repo = "sarin-jacob/clipboard"
Write-Host "Installing clipboard utility..."

# Formulate archive download link
$Url = "https://github.com/$Repo/releases/latest/download/clipboard_Windows_x86_64.zip"

# Create a temporary working directory
$TempDir = Join-Path $env:TEMP "clipboard_installer"
If (Test-Path $TempDir) { Remove-Item -Recurse -Force $TempDir }
New-Item -ItemType Directory -Path $TempDir | Out-Null

$ZipPath = Join-Path $TempDir "clipboard.zip"

Write-Host "Downloading archive from GitHub..."
Invoke-WebRequest -Uri $Url -OutFile $ZipPath

Write-Host "Extracting binary..."
Expand-Archive -Path $ZipPath -DestinationPath $TempDir -Force

# Execute the self-installation engine
Write-Host "Running system configuration engine..."
$BinaryPath = Join-Path $TempDir "clipboard.exe"
& $BinaryPath setup

# Clean up
Remove-Item -Recurse -Force $TempDir