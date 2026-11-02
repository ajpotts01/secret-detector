package config

import "regexp"

type Rule struct {
	ID          string
	Description string
	Pattern     *regexp.Regexp
}

const awsKeyPattern = `AKIA[0-9A-Z]{16}`
const gcpApiKeyPattern = `AIza[0-9A-Za-z\-_]{35}`
const githubPatPattern = `gh[pousr]_[A-Za-z0-9_]{36}`

func (r Rule) RunRule(line string) string {
	return r.Pattern.FindString(line)
}

func DefaultRules() []Rule {
	return []Rule{
		{
			ID:          "aws-key",
			Description: "AWS access keys",
			Pattern:     regexp.MustCompile(awsKeyPattern),
		},
		{
			ID: "gh-pat",
			Description: "Github personal access tokens",
			Pattern: regexp.MustCompile(githubPatPattern),
		},
		{
			ID: "gcp-key",
			Description: "Google Cloud API keys",
			Pattern: regexp.MustCompile(gcpApiKeyPattern),
		},
	}
}

func DefaultIgnoredDirs() []string {
	return []string{".git", "node_modules", "vendor", ".terraform", ".venv"}
}
