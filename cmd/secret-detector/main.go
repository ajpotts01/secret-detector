package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func start(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("secret-detector", flag.ContinueOnError)
	fs.SetOutput(stderr)

	path := fs.String("path", "", "path of file to scan")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *path == "" {
		return fmt.Errorf("-path is required")
	}

	fmt.Fprintf(stdout, "Scanning: %s\n", *path)
	return nil
}

func main() {
	if err := start(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) // is there a const for this?
	}
}