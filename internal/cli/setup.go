package cli

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func RunSetup() {
	// 1. Locate the currently running binary
	exePath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to determine executable path: %v\n", err)
		os.Exit(1)
	}
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		realPath = exePath
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not find user home directory: %v\n", err)
		os.Exit(1)
	}

	// 2. Define the unified user-local installation directory (~/.local/bin)
	targetDir := filepath.Join(homeDir, ".local", "bin")
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create target directory: %v\n", err)
		os.Exit(1)
	}

	binaryName := "clipboard"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	finalBinaryPath := filepath.Join(targetDir, binaryName)

	// 3. Move/Copy the binary to the destination directory if it isn't already there
	if realPath != finalBinaryPath {
		fmt.Printf("Copying binary to: %s\n", finalBinaryPath)
		if err := copyFile(realPath, finalBinaryPath); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to copy binary to central directory: %v\n", err)
			os.Exit(1)
		}
		// Make sure it remains executable on POSIX systems
		if runtime.GOOS != "windows" {
			_ = os.Chmod(finalBinaryPath, 0755)
		}
	}

	// 4. Handle Platform Specific Symlink & PATH Configuration
	if runtime.GOOS == "windows" {
		setupWindowsEnvironment(finalBinaryPath, targetDir)
	} else {
		setupUnixEnvironment(finalBinaryPath, targetDir)
	}
}

// setupUnixEnvironment handles symlinks and tracks path appending for Linux / macOS
func setupUnixEnvironment(binaryPath, targetDir string) {
	createShortcut(binaryPath, filepath.Join(targetDir, "cb"))
	createShortcut(binaryPath, filepath.Join(targetDir, "cv"))

	// Check if targetDir is in the active PATH environment variable
	currentPath := os.Getenv("PATH")
	if !strings.Contains(currentPath, targetDir) {
		fmt.Println("\nDetected that ~/.local/bin is not in your system $PATH.")
		
		shell := filepath.Base(os.Getenv("SHELL"))
		homeDir, _ := os.UserHomeDir()
		var rcFile string

		switch shell {
		case "zsh":
			rcFile = filepath.Join(homeDir, ".zshrc")
		case "bash":
			rcFile = filepath.Join(homeDir, ".bashrc")
		default:
			rcFile = filepath.Join(homeDir, ".profile")
		}

		exportLine := fmt.Sprintf("\n# Clipboard Utility Path Configuration\nexport PATH=\"$HOME/.local/bin:$PATH\"\n")
		
		f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Failed to automatically update your %s. Please add ~/.local/bin manually.\n", rcFile)
			return
		}
		defer f.Close()
		
		_, _ = f.WriteString(exportLine)
		fmt.Printf("Path update configuration injected into: %s\n", rcFile)
		fmt.Println("Please refresh your terminal environment: source " + rcFile)
	} else {
		fmt.Println("\nSetup complete! 'cb' and 'cv' commands are fully active and in your PATH.")
	}
}

// setupWindowsEnvironment configures PowerShell profile persistent aliases and user environment variable tracking
func setupWindowsEnvironment(binaryPath, targetDir string) {
	fmt.Println("Creating native Windows shim scripts...")

	exeName := filepath.Base(binaryPath)

	// 1. Generate Command Prompt / PowerShell Shims (.cmd)
	cbPath := filepath.Join(targetDir, "cb.cmd")
	cbContent := fmt.Sprintf("@echo off\r\n\"%%~dp0%s\" copy %%*\r\n", exeName)
	if err := os.WriteFile(cbPath, []byte(cbContent), 0755); err != nil {
		fmt.Printf("Failed to create cb.cmd: %v\n", err)
	} else {
		fmt.Println("Created cb.cmd")
	}

	cvPath := filepath.Join(targetDir, "cv.cmd")
	cvContent := fmt.Sprintf("@echo off\r\n\"%%~dp0%s\" paste %%*\r\n", exeName)
	if err := os.WriteFile(cvPath, []byte(cvContent), 0755); err != nil {
		fmt.Printf("Failed to create cv.cmd: %v\n", err)
	} else {
		fmt.Println("Created cv.cmd")
	}

	// 2. Generate Unix/Git Bash Shims (No extension)
	cbBash := filepath.Join(targetDir, "cb")
	cbBashContent := fmt.Sprintf("#!/bin/sh\nexec \"$(dirname \"$0\")/%s\" copy \"$@\"\n", exeName)
	_ = os.WriteFile(cbBash, []byte(cbBashContent), 0755)

	cvBash := filepath.Join(targetDir, "cv")
	cvBashContent := fmt.Sprintf("#!/bin/sh\nexec \"$(dirname \"$0\")/%s\" paste \"$@\"\n", exeName)
	_ = os.WriteFile(cvBash, []byte(cvBashContent), 0755)

	// 3. Ensure the folder is in the system PATH
	userPath, _ := os.LookupEnv("PATH")
	if !strings.Contains(strings.ToLower(userPath), strings.ToLower(targetDir)) {
		fmt.Println("Updating environment path definitions...")
		cmdText := fmt.Sprintf("[Environment]::SetEnvironmentVariable('Path', [Environment]::GetEnvironmentVariable('Path', 'User') + ';%s', 'User')", targetDir)
		_ = exec.Command("powershell", "-NoProfile", "-Command", cmdText).Run()
		fmt.Println("Added ~/.local/bin path to your Windows User Environment Path variables.")
	} else {
		fmt.Println("~/.local/bin is already in your Windows User Environment Path.")
	}

	fmt.Println("\nSetup complete! Please restart your terminal window.")
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func createShortcut(src, dst string) {
	_ = os.Remove(dst)
	err := os.Symlink(src, dst)
	if err!= nil {
		fmt.Printf("Failed to create symlink %s: %v\n", dst, err)
		return
	}
	fmt.Printf("Created symlink: %s -> %s\n", filepath.Base(dst), filepath.Base(src))
}