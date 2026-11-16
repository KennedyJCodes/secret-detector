package services

import "regexp"

// AWSAccessKeyRule detects AWS access key IDs. The recognized prefixes cover
// common IAM user and temporary, role, and service credentials.
type AWSAccessKeyRule struct{}

var awsAccessKeyID = regexp.MustCompile(`\b(?:AKIA|ASIA|AIDA|AROA|AIPA|ANPA|ANVA|ASCA)[A-Z0-9]{16}\b`)

// Name identifies this rule in scan results.
func (AWSAccessKeyRule) Name() string {
	return "AWS access key"
}

// Match reports whether text contains an AWS access key ID.
func (AWSAccessKeyRule) Match(text string) bool {
	return awsAccessKeyID.MatchString(text)
}
