package aws

import (
	"regexp"

	"github.com/KennedyJCodes/secret-detector/internal/rules"
)

var accessKeyIDPattern = regexp.MustCompile(`\b(?:AKIA|ASIA|AIDA|AROA|AIPA|ANPA|ANVA|ASCA)[A-Z0-9]{16}\b`)

type accessKeyIDRule struct{ rules.Rule }

func (r accessKeyIDRule) Match(text string) bool {
	for _, location := range accessKeyIDPattern.FindAllStringIndex(text, -1) {
		value := text[location[0]+4 : location[1]] // Skip the credential type prefix.
		if !isExampleKey(value) && hasAWSKeyContext(text, location[0], location[1]) {
			return true
		}
	}
	return false
}

// AWSAccessKeyRule detects AWS access key IDs such as AKIA... and ASIA....
func AWSAccessKeyRule() accessKeyIDRule {
	return accessKeyIDRule{Rule: rules.Rule{
		ID:          "aws-access-key",
		Description: "AWS access key",
		Pattern:     accessKeyIDPattern,
	}}
}
