package backend

import (
	"io"
	"os/exec"
)

type TermuxBackend struct{}

func (b *TermuxBackend) Name() string {
	return "Termux"
}

func (b *TermuxBackend) Available() bool {
	_, err := exec.LookPath("termux-clipboard-set")
	return err == nil
}

func (b *TermuxBackend) Copy(r io.Reader) error {
	cmd := exec.Command("termux-clipboard-set")
	cmd.Stdin = r
	return cmd.Run()
}

func (b *TermuxBackend) Paste(w io.Writer) error {
	cmd := exec.Command("termux-clipboard-get")
	cmd.Stdout = w
	return cmd.Run()
}