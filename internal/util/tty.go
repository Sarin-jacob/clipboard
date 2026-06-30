package util

import "os"

// IsTerminal checks if the given file (usually os.Stdin or os.Stdout) is attached to an interactive terminal.
func IsTerminal(f *os.File) bool {
	fileInfo, err := f.Stat()
	if err != nil {
		return false
	}
	// If the ModeCharDevice bit is set, it's a terminal (character device)
	// and NOT a pipe or redirect.
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// IsSSH checks if the current session is running over SSH.
func IsSSH() bool {
	return HasAnyEnv("SSH_CONNECTION", "SSH_CLIENT", "SSH_TTY")
}