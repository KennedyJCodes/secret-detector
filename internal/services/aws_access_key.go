package services

import (
	"regexp"
	"strings"

	"github.com/KennedyJCodes/secret-detector/internal/rules"
)

var awsAccessKeyID = regexp.MustCompile(`\b(?:AKIA|ASIA|AIDA|AROA|AIPA|ANPA|ANVA|ASCA)[A-Z0-9]{16}\b`)

var awsExampleMarkers = []string{"EXAMPLE", "SAMPLE", "PLACEHOLDER", "DUMMY", "TEST"}

// awsAccessKeyRule applies AWS-specific false-positive filtering on top of the
// shared pattern-based rule.
type awsAccessKeyRule struct{ rules.Rule }

func (r awsAccessKeyRule) Match(text string) bool {
	for _, match := range awsAccessKeyID.FindAllString(text, -1) {
		value := match[4:] // The first four characters are the credential type.
		if isAWSExampleKey(value) {
			continue
		}
		return true
	}
	return false
}

func isAWSExampleKey(value string) bool {
	for _, marker := range awsExampleMarkers {
		if strings.Contains(value, marker) {
			return true
		}
	}
	// Reject uniform placeholder payloads such as XXXXXXXXXXXXXXXXX or 0000000000000000.
	for i := 1; i < len(value); i++ {
		if value[i] != value[0] {
			return false
		}
	}
	return true
}

// AWSAccessKeyRule returns the detector for AWS access key IDs. The recognized
// prefixes cover common IAM user and temporary, role, and service credentials.
func AWSAccessKeyRule() awsAccessKeyRule {
	return awsAccessKeyRule{Rule: rules.Rule{
		ID:          "aws-access-key",
		Description: "AWS access key",
		Pattern:     awsAccessKeyID,
	}}
}
