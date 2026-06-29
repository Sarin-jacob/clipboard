package backend

import (
	"fmt"
	"io"
	"os/exec"
)

type MacBackend struct{}

func (b *MacBackend) Name() string {
	return "macOS (pbcopy/pbpaste)"
}

func (b *MacBackend) Available() bool {
	_, errCopy := exec.LookPath("pbcopy")
	_, errPaste := exec.LookPath("pbpaste")
	return errCopy == nil && errPaste == nil
}

func (b *MacBackend) Copy(r io.Reader) error {
	if !b.Available() {
		return fmt.Errorf("pbcopy is not available")
	}

	cmd := exec.Command("pbcopy")
	cmd.Stdin = r

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pbcopy failed: %w", err)
	}
	return nil
}

func (b *MacBackend) Paste(w io.Writer) error {
	if !b.Available() {
		return fmt.Errorf("pbpaste is not available")
	}

	cmd := exec.Command("pbpaste")
	cmd.Stdout = w

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pbpaste failed: %w", err)
	}
	return nil
}