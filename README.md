
# 📋 Clipboard

A lightweight, dependency-free, zero-configuration command-line utility that abstracts away platform-specific clipboard commands behind one predictable interface.

Instead of remembering `xclip`, `wl-copy`, `pbcopy`, `clip.exe`, or handling raw OSC52 escape sequences over SSH, just use **`cb`** and **`cv`**.


[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
\-
![Mac](https://img.shields.io/badge/mac%20os-000000?logo=apple&logoColor=white) 
![Windows](https://img.shields.io/badge/Windows-0078D6?logo=windows&logoColor=white) 
![Linux](https://img.shields.io/badge/Linux-FCC624?logo=linux&logoColor=black)
---

## ⚡ Key Features

* **Zero Dependencies:** Written in pure Go (`os/exec`). No C bindings, no runtime overhead.
* **Intelligent Environment Detection:** Automatically senses your context (SSH, Wayland, X11, WSL, macOS, Windows, Termux) and applies the correct priority rules.
* **OSC52 Support:** Copy text seamlessly from a remote SSH server directly onto your local machine's clipboard.
* **Unix-Pipeline Native:** Fully composable with standard piping (`|`) and file redirections (`>`).
* **Self-Healing Diagnostics:** Built-in health check and backend toggling for easy troubleshooting.

---

## 🚀 Installation

### Linux & macOS (Universal)
```bash
curl -fsSL https://raw.githubusercontent.com/sarin-jacob/clipboard/main/install.sh | sh
```

### Windows (PowerShell)

```powershell
irm "https://raw.githubusercontent.com/sarin-jacob/clipboard/main/install.ps1" | iex
```

### Manual Build / Go Toolchain

If you prefer not to pipe an installation script directly into your shell, you can set it up manually:

1. **Download the binary:** Go to the [Releases](https://github.com/sarin-jacob/clipboard/releases) page and download the raw binary matching your Operating System and Architecture (e.g., `clipboard_Linux_x86_64` or `clipboard_Windows_x86_64.exe`).
2. **Make it executable** *(Linux/macOS only)*:
   ```bash
   chmod +x clipboard_Linux_x86_64
   ```

3. **Initialize the setup engine:** 
   ```bash
   # On Linux/macOS
   ./clipboard_Linux_x86_64 setup

   # On Windows (PowerShell)
   .\clipboard_Windows_x86_64 setup

   ```



*(Alternatively, if you have a Go environment configured, you can just run the following)*

```bash
go install github.com/sarin-jacob/clipboard/cmd/clipboard@latest clipboard setup
```

*(Running `setup` automatically moves the binary to `~/.local/bin`, establishes your `cb`/`cv` symlinks/aliases, and updates your `$PATH` if needed!)*

---

## 📖 Usage & Examples

Once installed, the utility responds based on how it is invoked (`argv[0]` detection).

### Copying (`cb`)

```bash
# Pipe text to clipboard
echo "Hello World" | cb

# Copy contents of a file
cb notes.txt

# Copy raw text arguments directly
cb This string will be copied straight to the clipboard

```

### Pasting (`cv`)

```bash
# Paste to stdout
cv

# Pipe clipboard content to other utilities
cv | grep "API_KEY"

# Redirect clipboard content to a file
cv > backup.txt

```

---

## 🔍 Diagnostics & Advanced Tweaks

If you are ever unsure which clipboard engine is being picked up by the tool, you can ask it directly:

```bash
clipboard doctor
```

To list all backends supported by your operating system vs. which one is currently running:

```bash
clipboard list
```

### Forced Backends

Need to override the auto-detection for testing or scripting? Pass the `--backend` flag to force an engine:

```bash
cb --backend x11 "Force X11"
cv --backend wayland
```

---

## 🛡️ Backend Selection Matrix

When you run `cb`, **Clipboard** evaluates your environment in this exact order:

1. **Explicit Override:** Honors any user-specified `--backend` flag.
2. **OSC52 (Remote Terminal):** Triggered if `SSH_CONNECTION` is detected and `stdout` is a terminal.
3. **Wayland:** Targets `wl-copy` if `WAYLAND_DISPLAY` is active.
4. **X11:** Targets `xclip` (or falls back to `xsel`) if `DISPLAY` is active.
5. **macOS Native:** Targets `pbcopy` / `pbpaste`.
6. **Windows Native:** Targets `clip.exe` / PowerShell `Get-Clipboard`.
7. **WSL Bridge:** Accesses the underlying host Windows clipboard from within Linux via `/mnt/c/.../clip.exe`.
8. **Termux:** Targets Android clipboard interfaces via `termux-clipboard-set`.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

