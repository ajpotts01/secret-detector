package report

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ajpotts01/secret-detector/internal/detector"
)

type Format int

const (
	StdOutFormat Format = iota
	JSONFormat
)

var reportFormats = map[Format]string{
	StdOutFormat: "stdout",
	JSONFormat:   "json",
}

func (rf Format) String() string {
	return reportFormats[rf]
}

func Write(w io.Writer, findings []detector.Finding, format Format) error {
	// This might be best served as its own output function
	switch format {
	case StdOutFormat:
		for _, f := range findings {
			if _, err := fmt.Fprintf(w, "%s:%d [%s] %s\n", f.FilePath, f.LineNumber, f.RuleID, f.RedactedSecret); err != nil {
				return err
			}
		}
	case JSONFormat:
		// want "[]" instead of "null" when printing empty findings
		if findings == nil {
			findings = []detector.Finding{}
		}
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		return e.Encode(findings)
	}
	return nil
}
