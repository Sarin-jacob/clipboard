package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func RunSetup() {
	// 1. Find the path of the currently running 'clipboard' binary
	exePath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to determine executable path: %v\n", err)
		os.Exit(1)
	}

	// Resolve any existing symlinks to find the real binary
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		realPath = exePath
	}

	binDir := filepath.Dir(realPath)
	
	createShortcut(realPath, filepath.Join(binDir, "cb"))
	createShortcut(realPath, filepath.Join(binDir, "cv"))

	fmt.Println("\nSetup complete. You can now use 'cb' and 'cv'.")
	fmt.Printf("Symlinks created in: %s\n", binDir)
}

func createShortcut(target, alias string) {
	// Remove existing symlink if it exists to prevent errors
	_ = os.Remove(alias)

	// Windows requires Developer Mode or Admin rights to create symlinks via os.Symlink.
	// As a fallback, we could write a small .bat wrapper, but for now, we warn the user.
	err := os.Symlink(target, alias)
	if err != nil {
		if runtime.GOOS == "windows" {
			fmt.Printf("Failed to create %s. On Windows, you may need to run as Administrator, or set up a PowerShell alias instead:\n", filepath.Base(alias))
			fmt.Printf("   Set-Alias -Name %s -Value %s\n", filepath.Base(alias), target)
			return
		}
		fmt.Printf("Failed to create symlink %s: %v\n", alias, err)
		return
	}
	
	fmt.Printf("Created symlink: %s -> %s\n", filepath.Base(alias), filepath.Base(target))
}