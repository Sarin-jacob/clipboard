package backend

import (
	"fmt"
	"io"
	"os/exec"
)

type X11Backend struct{}

func (b *X11Backend) Name() string {
	return "X11 (xclip)"
}

func (b *X11Backend) Available() bool {
	_, err := exec.LookPath("xclip")
	return err == nil
}

func (b *X11Backend) Copy(r io.Reader) error {
	if !b.Available() {
		return fmt.Errorf("xclip is not installed")
	}

	// xclip requires -selection clipboard to use the standard Ctrl+C/Ctrl+V clipboard
	// rather than the middle-mouse primary selection.
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = r

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("xclip failed: %w", err)
	}
	return nil
}

func (b *X11Backend) Paste(w io.Writer) error {
	if !b.Available() {
		return fmt.Errorf("xclip is not installed")
	}

	cmd := exec.Command("xclip", "-selection", "clipboard", "-o")
	cmd.Stdout = w

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("xclip paste failed: %w", err)
	}
	return nil
}