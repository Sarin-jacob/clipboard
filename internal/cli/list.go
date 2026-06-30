package cli

import (
	"fmt"
	"github.com/Sarin-jacob/clipboard/internal/backend"
	"github.com/Sarin-jacob/clipboard/internal/detect"
)

func RunList() {
	fmt.Println("Available\n")

	// Instantiate all backends to check their status
	backends := []backend.Backend{
		&backend.OSC52Backend{},
		&backend.WaylandBackend{},
		&backend.X11Backend{},
		&backend.MacBackend{},
		&backend.WindowsBackend{},
		&backend.WSLBackend{},
		&backend.TermuxBackend{},
	}

	for _, b := range backends {
		if b.Available() {
			fmt.Printf("✓ %s\n", b.Name())
		} else {
			fmt.Printf("✗ %s\n", b.Name())
		}
	}

	fmt.Println("\nSelected\n")
	
	selected, err := detect.GetPreferredBackend("")
	if err != nil {
		fmt.Println("None (No supported backend found)")
	} else {
		fmt.Println(selected.Name())
	}
}