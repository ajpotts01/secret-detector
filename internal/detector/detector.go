package detector

import (
	"github.com/ajpotts01/secret-detector/internal/config"
)

type Finding struct {
	FilePath       string
	LineNumber     int
	RuleID         string
	RedactedSecret string
}

type Engine struct {
	rules []config.Rule
}

func NewEngine(rules []config.Rule) *Engine {
	return &Engine{
		rules: rules,
	}
}

// InspectLine checks each rule's Pattern.FindString(line).
// If non-empty, append a Finding with a redacted copy of the match.
func (e *Engine) InspectLine(filePath string, lineNumber int, line string) []Finding {
	var findings []Finding
	for _, r := range e.rules {
		if match := r.RunRule(line); match != "" {
			// create finding
			f := Finding{
				RuleID:         r.ID,
				FilePath:       filePath,
				LineNumber:     lineNumber,
				RedactedSecret: redact(match),
			}
			findings = append(findings, f)
		}
	}

	return findings
}

// Redact masks the middle of a matched secret (e.g. "AKIA...XXXX")
func redact(secret string) string {
	if len(secret) < 4 {
		return "..."
	}
	if len(secret) <= 8 {
		return secret[:2] + "..."
	}
	return secret[:4] + "..." + secret[len(secret)-4:]
}
