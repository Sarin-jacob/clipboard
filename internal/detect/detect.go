package detect

import (
	"fmt"
	"strings"

	"github.com/Sarin-jacob/clipboard/internal/backend"
)

// GetPreferredBackend iterates through available backends based on your priority list.
func GetPreferredBackend(forcedBackend string) (backend.Backend, error) {
	// 1. If the user explicitly requested a backend via flag
	if forcedBackend != "" {
		return lookupByName(forcedBackend)
	}

	// 2. Define the priority order
	// This list matches your design document order.
	priorityList := []backend.Backend{
		&backend.OSC52Backend{},
		&backend.WaylandBackend{},
		&backend.X11Backend{},
		&backend.MacBackend{},
		&backend.WindowsBackend{},
		&backend.WSLBackend{},
		&backend.TermuxBackend{},
	}

	// 3. Iterate and find the first one that is Available()
	for _, b := range priorityList {
		if b.Available() {
			return b, nil
		}
	}

	return nil, fmt.Errorf("no supported clipboard backend found")
}

// lookupByName maps a CLI flag string to a backend instance
func lookupByName(name string) (backend.Backend, error) {
	switch strings.ToLower(name) {
	case "osc52":
		return &backend.OSC52Backend{}, nil
	case "wayland":
		return &backend.WaylandBackend{}, nil
	case "xclip", "x11":
		return &backend.X11Backend{}, nil
	case "mac", "pbcopy":
		return &backend.MacBackend{}, nil
	case "windows", "clip":
		return &backend.WindowsBackend{}, nil
	case "wsl":
		return &backend.WSLBackend{}, nil
	case "termux":
		return &backend.TermuxBackend{}, nil
	default:
		return nil, fmt.Errorf("unknown backend: %s", name)
	}
}