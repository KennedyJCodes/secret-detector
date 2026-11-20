package aws

import (
	"regexp"
	"strings"
)

var camelCaseBoundary = regexp.MustCompile(`([a-z0-9])([A-Z])`)

var exampleMarkers = []string{"EXAMPLE", "SAMPLE", "PLACEHOLDER", "DUMMY", "TEST"}

var keyContextTerms = map[string]struct{}{
	"access": {}, "credential": {}, "credentials": {}, "key": {}, "keys": {},
	"secret": {}, "secrets": {}, "token": {}, "tokens": {}, "aws": {},
}

const contextDistance = 48

func isExampleKey(value string) bool {
	for _, marker := range exampleMarkers {
		if strings.Contains(value, marker) {
			return true
		}
	}
	for i := 1; i < len(value); i++ {
		if value[i] != value[0] {
			return false
		}
	}
	return true
}

func hasAWSKeyContext(text string, keyStart, keyEnd int) bool {
	return hasNearbyContext(text, keyStart, keyEnd, hasKeyContextTerm)
}

func hasSecretKeyContext(text string, keyStart, keyEnd int) bool {
	return hasNearbyContext(text, keyStart, keyEnd, hasSecretKeyContextTerm)
}

func hasNearbyContext(text string, keyStart, keyEnd int, matches func(string) bool) bool {
	start := keyStart - contextDistance
	if start < 0 {
		start = 0
	}
	end := keyEnd + contextDistance
	if end > len(text) {
		end = len(text)
	}
	lineStart := strings.LastIndex(text[:keyStart], "\n") + 1
	if start < lineStart {
		start = lineStart
	}
	if lineEnd := strings.IndexByte(text[keyEnd:], '\n'); lineEnd >= 0 && keyEnd+lineEnd < end {
		end = keyEnd + lineEnd
	}
	return matches(text[start:end])
}

func hasKeyContextTerm(context string) bool {
	for _, word := range contextWords(context) {
		if _, ok := keyContextTerms[word]; ok {
			return true
		}
	}
	return false
}

func hasSecretKeyContextTerm(context string) bool {
	for _, word := range contextWords(context) {
		if word == "secret" || word == "secrets" || word == "credential" || word == "credentials" {
			return true
		}
	}
	return false
}

func contextWords(context string) []string {
	context = camelCaseBoundary.ReplaceAllString(context, "$1 $2")
	return strings.FieldsFunc(strings.ToLower(context), func(r rune) bool {
		return r < 'a' || r > 'z'
	})
}
