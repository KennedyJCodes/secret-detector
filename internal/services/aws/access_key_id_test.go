package aws

import "testing"

func TestAWSAccessKeyRuleMatchesSupportedPrefixes(t *testing.T) {
	for _, key := range []string{
		"AKIA1234567890ABCDEF", "ASIA1234567890ABCDEF", "AIDA1234567890ABCDEF", "AROA1234567890ABCDEF",
		"AIPA1234567890ABCDEF", "ANPA1234567890ABCDEF", "ANVA1234567890ABCDEF", "ASCA1234567890ABCDEF",
	} {
		if !AWSAccessKeyRule().Match("credential=" + key) {
			t.Errorf("AWSAccessKeyRule did not match %q", key)
		}
	}
}

func TestAWSAccessKeyRuleRejectsInvalidValues(t *testing.T) {
	for _, text := range []string{
		"AKIA1234567890ABCDE", "AKIA1234567890ABCDEFG", "akia1234567890abcdef", "ZZZZ1234567890ABCDEF", "",
	} {
		if AWSAccessKeyRule().Match(text) {
			t.Errorf("AWSAccessKeyRule unexpectedly matched %q", text)
		}
	}
}

func TestAWSAccessKeyRuleMetadata(t *testing.T) {
	rule := AWSAccessKeyRule()
	if rule.ID != "aws-access-key" || rule.Name() != "AWS access key" {
		t.Errorf("unexpected rule metadata: ID=%q name=%q", rule.ID, rule.Name())
	}
}
