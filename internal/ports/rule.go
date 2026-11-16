package ports

// Rule is the common contract implemented by every secret detection rule.
// Implementations should be safe to reuse across lines and scans.
type Rule interface {
	Name() string
	Match(text string) bool
}
