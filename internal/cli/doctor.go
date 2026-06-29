package cli

import (
	"fmt"

	"github.com/Sarin-jacob/clipboard/internal/backend"
	"github.com/Sarin-jacob/clipboard/internal/detect"
)

func RunDoctor() {
	fmt.Println("--- Clipboard Diagnostics ---")
	
	// 1. Check OS/Env
	fmt.Printf("OS: %s\n", "Linux/macOS/Windows") // Can use runtime.GOOS
	
	// 2. Check Backends
	list := []backend.Backend{
		&backend.OSC52Backend{},
		&backend.WaylandBackend{},
		&backend.X11Backend{},
		&backend.MacBackend{},
		&backend.WindowsBackend{},
	}
	
	for _, b := range list {
		status := "✗"
		if b.Available() {
			status = "✓"
		}
		fmt.Printf("%s %s\n", status, b.Name())
	}
    
    selected, _ := detect.GetPreferredBackend("")
    if selected != nil {
        fmt.Printf("\nCurrently selected backend: %s\n", selected.Name())
    }
}