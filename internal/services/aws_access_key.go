package services

import (
	"regexp"

	"github.com/KennedyJCodes/secret-detector/internal/rules"
)

var awsAccessKeyID = regexp.MustCompile(`\b(?:AKIA|ASIA|AIDA|AROA|AIPA|ANPA|ANVA|ASCA)[A-Z0-9]{16}\b`)

// AWSAccessKeyRule returns the detector for AWS access key IDs. The recognized
// prefixes cover common IAM user and temporary, role, and service credentials.
func AWSAccessKeyRule() rules.Rule {
	return rules.Rule{
		ID:          "aws-access-key",
		Description: "AWS access key",
		Pattern:     awsAccessKeyID,
	}
}
