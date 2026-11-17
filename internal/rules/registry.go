// Package rules provides the shared registry used to configure and apply rules.
package rules

import (
	"fmt"
	"math"
	"regexp"

	"github.com/KennedyJCodes/secret-detector/internal/ports"
)

// Rule describes a pattern-based detector. MinEntropy is optional; zero means
// that matching depends on the pattern alone.
type Rule struct {
	ID          string
	Description string
	Pattern     *regexp.Regexp
	MinEntropy  float64
}

// Name identifies the rule in scan results.
func (r Rule) Name() string { return r.Description }

// Match reports whether text contains a match that satisfies the rule.
func (r Rule) Match(text string) bool {
	if r.Pattern == nil {
		return false
	}
	match := r.Pattern.FindString(text)
	if match == "" {
		return false
	}
	return r.MinEntropy <= 0 || shannonEntropy(match) >= r.MinEntropy
}

// Registry holds the rules enabled for a scan.
type Registry struct {
	rules []ports.Rule
}

// NewRegistry creates an empty registry. Rules can be added with Add.
func NewRegistry() *Registry {
	return &Registry{}
}

// Add validates and registers a rule.
func (r *Registry) Add(rule ports.Rule) error {
	if r == nil {
		return fmt.Errorf("rule registry cannot be nil")
	}
	if rule == nil {
		return fmt.Errorf("rule cannot be nil")
	}
	if definition, ok := rule.(Rule); ok {
		if definition.ID == "" {
			return fmt.Errorf("rule ID cannot be empty")
		}
		if definition.Description == "" {
			return fmt.Errorf("rule description cannot be empty")
		}
		if definition.Pattern == nil {
			return fmt.Errorf("rule pattern cannot be nil")
		}
		if math.IsNaN(definition.MinEntropy) || definition.MinEntropy < 0 {
			return fmt.Errorf("rule minimum entropy must be zero or greater")
		}
	}
	r.rules = append(r.rules, rule)
	return nil
}

// Matches returns the names of all rules that match text, in registration order.
func (r *Registry) Matches(text string) []string {
	if r == nil {
		return nil
	}
	var matches []string
	for _, rule := range r.rules {
		if rule.Match(text) {
			matches = append(matches, rule.Name())
		}
	}
	return matches
}

func shannonEntropy(value string) float64 {
	if len(value) == 0 {
		return 0
	}
	var counts [256]int
	for i := 0; i < len(value); i++ {
		counts[value[i]]++
	}
	entropy := 0.0
	for _, count := range counts {
		if count == 0 {
			continue
		}
		probability := float64(count) / float64(len(value))
		entropy -= probability * math.Log2(probability)
	}
	return entropy
}

var _ ports.Rule = Rule{}
