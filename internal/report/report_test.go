package report_test

import (
	"bytes"
	"testing"

	"github.com/ajpotts01/secret-detector/internal/detector"
	"github.com/ajpotts01/secret-detector/internal/report"
)

func TestReport(t *testing.T) {
	findings := []detector.Finding{
		{
			FilePath:       "test.txt",
			LineNumber:     24,
			RuleID:         "aws-key",
			RedactedSecret: "AKIA...MPLE",
		},
	}

	var buf bytes.Buffer
	want := "test.txt:24 [aws-key] AKIA...MPLE\n"
	err := report.Write(&buf, findings, report.StdOutFormat)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	got := buf.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
