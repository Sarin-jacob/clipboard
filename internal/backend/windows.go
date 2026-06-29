package backend

import (
	"fmt"
	"io"
	"os/exec"
)

type WindowsBackend struct{}

func (b *WindowsBackend) Name() string {
	return "Windows (clip.exe)"
}

func (b *WindowsBackend) Available() bool {
	// clip.exe is in the system path by default on Windows
	_, err := exec.LookPath("clip.exe")
	return err == nil
}

func (b *WindowsBackend) Copy(r io.Reader) error {
	if !b.Available() {
		return fmt.Errorf("clip.exe not found")
	}

	cmd := exec.Command("clip.exe")
	cmd.Stdin = r
	return cmd.Run()
}

func (b *WindowsBackend) Paste(w io.Writer) error {
	// PowerShell Get-Clipboard is the standard way to retrieve content
	cmd := exec.Command("powershell", "-Command", "Get-Clipboard")
	cmd.Stdout = w
	return cmd.Run()
}