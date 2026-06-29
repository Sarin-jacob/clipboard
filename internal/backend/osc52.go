package backend

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

type OSC52Backend struct{}

func (b *OSC52Backend) Name() string {
	return "OSC52 (Remote Terminal)"
}

func (b *OSC52Backend) Available() bool {
	// Simple check: Is stdout a terminal? 
	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func (b *OSC52Backend) Copy(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// Format: \x1b]52;c;{base64_data}\x07
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Fprintf(os.Stdout, "\x1b]52;c;%s\x07", encoded)
	return nil
}

func (b *OSC52Backend) Paste(w io.Writer) error {
	return fmt.Errorf("OSC52 does not support pasting (read-only protocol)")
}