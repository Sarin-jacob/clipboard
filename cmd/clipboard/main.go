package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/Sarin-jacob/clipboard/internal/cli"
	"github.com/Sarin-jacob/clipboard/internal/util"
)

func printUsage() {
	fmt.Println("Clipboard - A Cross-Platform Clipboard Utility")
	fmt.Println("\nUsage via Symlinks (Recommended):")
	fmt.Println("  cb [text|file]        Copy text, file contents, or stdin to clipboard")
	fmt.Println("  cv                    Output clipboard contents to stdout")
	
	fmt.Println("\nUsage via Binary:")
	fmt.Println("  clipboard copy [text] Copy text to clipboard")
	fmt.Println("  clipboard paste       Output clipboard contents")
	fmt.Println("  clipboard list        List available clipboard backends")
	fmt.Println("  clipboard doctor      Run diagnostic health check")
	fmt.Println("  clipboard setup       Automatically generate cb/cv symlinks")

	fmt.Println("\nExamples:")
	fmt.Println("  echo \"hello\" | cb")
	fmt.Println("  cb notes.txt")
	fmt.Println("  cv > backup.txt")
}

func main() {
	name := filepath.Base(os.Args[0])

	// Strip .exe extension on Windows to normalize comparison
	if ext := filepath.Ext(name); ext != "" {
		name = name[:len(name)-len(ext)]
	}

	switch name {
	case "cb":
		cli.RunCopy(os.Args[1:])
	case "cv":
		cli.RunPaste(os.Args[1:])
	default:
		// If running as "clipboard", check if data is being piped into it directly.
		// If stdin is a pipe (not a terminal) and no command was given, auto-run copy!
		if !util.IsTerminal(os.Stdin) && len(os.Args) < 2 {
			cli.RunCopy(os.Args[1:])
			return
		}

		if len(os.Args) < 2 {
			printUsage()
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
		case "setup":
			cli.RunSetup()
		default:
			fmt.Printf("Unknown command: %s\n\n", os.Args[1])
			printUsage()
			os.Exit(1)
		}
	}
}