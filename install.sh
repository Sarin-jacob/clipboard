#!/bin/sh
set -e

REPO="sarin-jacob/clipboard"
echo "Installing clipboard utility..."

# 1. Detect Environment
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64)  ARCH="x86_64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
    linux)   OS="Linux" ;;
    darwin)  OS="Darwin" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# 2. Formulate Asset Download URL
URL="https://github.com/$REPO/releases/latest/download/clipboard_${OS}_${ARCH}"

# 3. Download directly to a temporary location
TEMP_BIN=$(mktemp)
echo "Downloading raw binary from GitHub..."
curl -fsSL "$URL" -o "$TEMP_BIN"
chmod +x "$TEMP_BIN"

# 4. Let the Go binary handle its own relocation and environment setup!
echo "Running system configuration engine..."
"$TEMP_BIN" setup

# Clean up the temp file
rm "$TEMP_BIN"