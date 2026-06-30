package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/Sarin-jacob/clipboard/internal/detect"
)

func RunCopy(args []string) {
	fs := flag.NewFlagSet("copy", flag.ExitOnError)
	backendFlag := fs.String("backend", "", "Force specific backend")
	fs.Parse(args)

	b, err := detect.GetPreferredBackend(*backendFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var input io.Reader

	// Check if args were passed after flags
	if fs.NArg() > 0 {
		target := fs.Arg(0)
		
		// 1. Check if it's an existing file on disk
		if stat, err := os.Stat(target); err == nil && !stat.IsDir() {
			file, err := os.Open(target)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()
			input = file
		} else {
			// 2. If it's not a file, join all remaining arguments as raw text
			rawText := strings.Join(fs.Args(), " ")
			input = strings.NewReader(rawText)
		}
	} else {
		// 3. Fall back to standard input piping (e.g., echo "hi" | cb)
		input = os.Stdin
	}

	if err := b.Copy(input); err != nil {
		fmt.Fprintf(os.Stderr, "Copy failed: %v\n", err)
		os.Exit(1)
	}
}