package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ajpotts01/secret-detector/internal/detector"
	"github.com/ajpotts01/secret-detector/internal/input"
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

	var findings []string
	err := input.ReadInputSingle(os.DirFS("."), *path, func(lineNum int, line string) {
		newFindings := detector.ScanLine(lineNum, line)
		findings = append(findings, newFindings...)
	})
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	fmt.Fprintln(stdout, findings)
	return nil
}

func main() {
	if err := start(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) // is there a const for this?
	}
}
