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
URL="https://github.com/$REPO/releases/latest/download/clipboard_${OS}_${ARCH}.tar.gz"

# 3. Create a secure temporary directory
TEMP_DIR=$(mktemp -d)
TAR_PATH="$TEMP_DIR/clipboard.tar.gz"

echo "Downloading archive from GitHub..."
curl -fsSL "$URL" -o "$TAR_PATH"

echo "Extracting binary..."
tar -xzf "$TAR_PATH" -C "$TEMP_DIR"

# 4. Execute the self-installation engine
echo "Running system configuration engine..."
"$TEMP_DIR/clipboard" setup

# 5. Clean up
rm -rf "$TEMP_DIR"