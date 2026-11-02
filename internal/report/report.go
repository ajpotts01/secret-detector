package report

import (
	"fmt"
	"io"

	"github.com/ajpotts01/secret-detector/internal/detector"
)

func Write(w io.Writer, findings []detector.Finding) error {
	for _, f := range findings {
		if _, err := fmt.Fprintf(w, "%s:%d [%s] %s\n", f.FilePath, f.LineNumber, f.RuleID, f.RedactedSecret); err != nil {
			return err
		}
	}

	return nil
}
