package backend

import "io"

// Backend represents a system clipboard implementation.
type Backend interface {
	// Name returns the display name of the backend (e.g., "Wayland (wl-copy)")
	Name() string

	// Available checks if the required underlying executables exist in the PATH
	Available() bool

	// Copy reads from the provided reader and writes to the system clipboard
	Copy(r io.Reader) error

	// Paste reads from the system clipboard and writes to the provided writer
	Paste(w io.Writer) error
}