package ports

// Rule is the common contract implemented by every secret detection rule.
// Implementations should be safe to reuse across lines and scans.
type Rule interface {
	Name() string
	Match(text string) bool
}

// RuleSet applies the configured rules to text and returns the names of rules
// that matched. Implementations should return matches in a stable order.
type RuleSet interface {
	Matches(text string) []string
}
