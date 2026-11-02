package config

import "regexp"

type Rule struct {
	ID          string
	Description string
	Pattern     *regexp.Regexp
}

const awsKeyPattern = `AKIA[0-9A-Z]{16}`

func (r Rule) RunRule(line string) string {
	return r.Pattern.FindString(line)
}

func DefaultRules() []Rule {
	// Return 2-3 basic rules using regexp.MustCompile(...)
	// e.g. AWS Access Key: `AKIA[0-9A-Z]{16}`
	return []Rule{
		{
			ID:          "aws-key",
			Description: "AWS access keys",
			Pattern:     regexp.MustCompile(awsKeyPattern),
		},
	}
}

func DefaultIgnoredDirs() []string {
	// Return []string{".git", "node_modules", "vendor", ".terraform"}
	return []string{".git", "node_modules", "vendor", ".terraform", ".venv"}
}
