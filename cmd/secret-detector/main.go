package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ajpotts01/secret-detector/internal/config"
	"github.com/ajpotts01/secret-detector/internal/detector"
	"github.com/ajpotts01/secret-detector/internal/input"
	"github.com/ajpotts01/secret-detector/internal/report"
)

var ErrHasFindings = errors.New("secrets detected")

func start(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("secret-detector", flag.ContinueOnError)
	fs.SetOutput(stderr)

	path := fs.String("path", "", "path of file to scan")
	format := fs.String("format", "stdout", "format to output findings - stdout or json")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *path == "" {
		return fmt.Errorf("-path is required")
	}

	var rf report.Format
	switch *format {
	case "stdout":
		rf = report.StdOutFormat
	case "json":
		rf = report.JSONFormat
	default:
		return fmt.Errorf("incorrect format: must specify json or stdout")
	}

	fmt.Fprintf(stderr, "Scanning: %s\n", *path) //nolint:errcheck

	var findings []detector.Finding
	ruleEngine := detector.NewEngine(config.DefaultRules())
	err := input.ReadDir(os.DirFS("."), *path, func(filePath string, lineNum int, line string) {
		newFindings := ruleEngine.InspectLine(filePath, lineNum, line)
		findings = append(findings, newFindings...)
	})
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	if err := report.Write(stdout, findings, rf); err != nil {
		return err
	}
	if len(findings) > 0 {
		return ErrHasFindings
	}
	return nil
}

func main() {
	if err := start(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) // is there a const for this?
	}
}
