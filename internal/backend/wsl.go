package backend

import (
	"io"
	"os/exec"
)

type WSLBackend struct{}

func (b *WSLBackend) Name() string {
	return "WSL (Windows Bridge)"
}

func (b *WSLBackend) Available() bool {
	// Check if the WSL interop path exists
	_, err := exec.LookPath("/mnt/c/Windows/System32/clip.exe")
	return err == nil
}

func (b *WSLBackend) Copy(r io.Reader) error {
	cmd := exec.Command("/mnt/c/Windows/System32/clip.exe")
	cmd.Stdin = r
	return cmd.Run()
}

func (b *WSLBackend) Paste(w io.Writer) error {
	// Use PowerShell to get clipboard content
	cmd := exec.Command("/mnt/c/Windows/System32/WindowsPowerShell/v1.0/powershell.exe", "-Command", "Get-Clipboard")
	cmd.Stdout = w
	return cmd.Run()
}