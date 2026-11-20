package aws

import (
	"regexp"
	"strings"

	"github.com/KennedyJCodes/secret-detector/internal/rules"
)

var secretAccessKeyPattern = regexp.MustCompile(`[A-Za-z0-9/+=]{40}`)

type secretAccessKeyRule struct{ rules.Rule }

func (r secretAccessKeyRule) Match(text string) bool {
	for _, location := range secretAccessKeyPattern.FindAllStringIndex(text, -1) {
		value := text[location[0]:location[1]]
		if hasSecretKeyBoundary(text, location[0], location[1]) &&
			!isExampleKey(value) && hasSecretKeyContext(text, location[0], location[1]) {
			return true
		}
	}
	return false
}

// AWSSecretAccessKeyRule detects AWS secret access key values near a secret label.
func AWSSecretAccessKeyRule() secretAccessKeyRule {
	return secretAccessKeyRule{Rule: rules.Rule{
		ID:          "aws-secret-access-key",
		Description: "AWS secret access key",
		Pattern:     secretAccessKeyPattern,
	}}
}

func hasSecretKeyBoundary(text string, start, end int) bool {
	const keyCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789/+="
	return (start == 0 || !strings.ContainsRune(keyCharacters, rune(text[start-1]))) &&
		(end == len(text) || !strings.ContainsRune(keyCharacters, rune(text[end])))
}
