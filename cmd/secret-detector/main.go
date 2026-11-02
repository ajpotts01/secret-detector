package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ajpotts01/secret-detector/internal/config"
	"github.com/ajpotts01/secret-detector/internal/detector"
	"github.com/ajpotts01/secret-detector/internal/input"
	"github.com/ajpotts01/secret-detector/internal/report"
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

	fmt.Fprintf(stdout, "Scanning: %s\n", *path) //nolint:errcheck

	var findings []detector.Finding
	ruleEngine := detector.NewEngine(config.DefaultRules())
	err := input.ReadDir(os.DirFS("."), *path, func(filePath string, lineNum int, line string) {
		newFindings := ruleEngine.InspectLine(filePath, lineNum, line)
		findings = append(findings, newFindings...)
	})
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return report.Write(stdout, findings)
}

func main() {
	if err := start(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) // is there a const for this?
	}
}
