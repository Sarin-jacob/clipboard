package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	
	"github.com/Sarin-jacob/clipboard/internal/detect"
)

func RunCopy(args []string) {
	fs := flag.NewFlagSet("copy", flag.ExitOnError)
	backendFlag := fs.String("backend", "", "Force specific backend")
	fs.Parse(args)

	// Determine backend
	b, err := detect.GetPreferredBackend(*backendFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var input io.Reader = os.Stdin
	
	// If a filename is provided as an argument, use that instead of stdin
	if fs.NArg() > 0 {
		file, err := os.Open(fs.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	if err := b.Copy(input); err != nil {
		fmt.Fprintf(os.Stderr, "Copy failed: %v\n", err)
		os.Exit(1)
	}
}