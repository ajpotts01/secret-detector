package detector

import (
	"fmt"
	"regexp"
)

const AwsKeyPattern = `AKIA[0-9A-Z]{16}`

var awsKey = regexp.MustCompile(AwsKeyPattern)

func ScanLine(lineNum int, line string) []string {
	var findings []string
	matched := findAWSKey(line)

	if matched {
		findings = append(findings, fmt.Sprintf("Found AWS key on line %d", lineNum))
	}
	return findings
}

func findAWSKey(line string) bool {
	return awsKey.MatchString(line)
}
