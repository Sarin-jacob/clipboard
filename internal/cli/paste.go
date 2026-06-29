package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sarin-jacob/clipboard/internal/detect"
)

func RunPaste(args []string) {
	fs := flag.NewFlagSet("paste", flag.ExitOnError)
	backendFlag := fs.String("backend", "", "Force specific backend")
	fs.Parse(args)

	b, err := detect.GetPreferredBackend(*backendFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := b.Paste(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Paste failed: %v\n", err)
		os.Exit(1)
	}
}