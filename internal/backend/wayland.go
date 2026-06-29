package backend

import (
	"fmt"
	"io"
	"os/exec"
)

type WaylandBackend struct{}

func (b *WaylandBackend) Name() string {
	return "Wayland (wl-clipboard)"
}

func (b *WaylandBackend) Available() bool {
	_, errCopy := exec.LookPath("wl-copy")
	_, errPaste := exec.LookPath("wl-paste")
	return errCopy == nil && errPaste == nil
}

func (b *WaylandBackend) Copy(r io.Reader) error {
	if !b.Available() {
		return fmt.Errorf("wl-copy is not installed")
	}

	cmd := exec.Command("wl-copy")
	cmd.Stdin = r
	return cmd.Run()
}

func (b *WaylandBackend) Paste(w io.Writer) error {
	if !b.Available() {
		return fmt.Errorf("wl-paste is not installed")
	}

	cmd := exec.Command("wl-paste")
	cmd.Stdout = w
	return cmd.Run()
}