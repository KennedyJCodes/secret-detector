package services

import "testing"

func TestAWSAccessKeyRuleMatchesSupportedPrefixes(t *testing.T) {
	for _, key := range []string{
		"AKIA1234567890ABCDEF",
		"ASIA1234567890ABCDEF",
		"AIDA1234567890ABCDEF",
		"AROA1234567890ABCDEF",
		"AIPA1234567890ABCDEF",
		"ANPA1234567890ABCDEF",
		"ANVA1234567890ABCDEF",
		"ASCA1234567890ABCDEF",
	} {
		if !AWSAccessKeyRule().Match("credential=" + key) {
			t.Errorf("AWSAccessKeyRule did not match %q", key)
		}
	}
}

func TestAWSAccessKeyRuleRejectsInvalidValues(t *testing.T) {
	rule := AWSAccessKeyRule()
	for _, text := range []string{
		"AKIA1234567890ABCDE",   // too short (19)
		"AKIA1234567890ABCDEFG", // too long (21)
		"akia1234567890abcdef",  // lowercase prefix
		"ZZZZ1234567890ABCDEF",  // unsupported prefix
		"",
	} {
		if rule.Match(text) {
			t.Errorf("AWSAccessKeyRule unexpectedly matched %q", text)
		}
	}
}

func TestAWSAccessKeyRuleMetadata(t *testing.T) {
	rule := AWSAccessKeyRule()
	if rule.ID != "aws-access-key" {
		t.Errorf("rule ID = %q, want aws-access-key", rule.ID)
	}
	if rule.Name() != "AWS access key" {
		t.Errorf("rule name = %q, want AWS access key", rule.Name())
	}
}
