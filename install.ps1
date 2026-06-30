# install.ps1
$Repo = "sarin-jacob/clipboard"
Write-Host "Installing clipboard utility..."

# Formulate binary download link
$Url = "https://github.com/$Repo/releases/latest/download/clipboard_Windows_x86_64"

$TempExe = Join-Path $env:TEMP "clipboard_installer.exe"

Write-Host "Downloading raw binary from GitHub..."
Invoke-WebRequest -Uri $Url -OutFile $TempExe

Write-Host "Running system configuration engine..."
# Execute the downloaded file to let it install itself
& $TempExe setup

# Clean up temp installer
Remove-Item $TempExe