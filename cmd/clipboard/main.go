package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/Sarin-jacob/clipboard/internal/cli"
)

func main() {
	// Determine mode based on executable name
	name := filepath.Base(os.Args[0])

	switch name {
	case "cb":
		cli.RunCopy(os.Args[1:])
	case "cv":
		cli.RunPaste(os.Args[1:])
	default:
		// Fallback for "clipboard" command or others
		if len(os.Args) < 2 {
			fmt.Println("Usage: clipboard [copy|paste|list|doctor]")
			os.Exit(1)
		}
		
		switch os.Args[1] {
		case "copy", "cb":
			cli.RunCopy(os.Args[2:])
		case "paste", "cv":
			cli.RunPaste(os.Args[2:])
		case "list":
			cli.RunList()
		case "doctor":
			cli.RunDoctor()
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
			os.Exit(1)
		}
	}
}