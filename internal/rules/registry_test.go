package rules

import (
	"math"
	"regexp"
	"testing"
)

func TestRegistryAddValidatesRuleDefinitions(t *testing.T) {
	tests := []struct {
		name string
		rule Rule
	}{
		{name: "empty ID", rule: Rule{Description: "rule", Pattern: regexp.MustCompile("x")}},
		{name: "empty description", rule: Rule{ID: "id", Pattern: regexp.MustCompile("x")}},
		{name: "nil pattern", rule: Rule{ID: "id", Description: "rule"}},
		{name: "negative entropy", rule: Rule{ID: "id", Description: "rule", Pattern: regexp.MustCompile("x"), MinEntropy: -1}},
		{name: "NaN entropy", rule: Rule{ID: "id", Description: "rule", Pattern: regexp.MustCompile("x"), MinEntropy: math.NaN()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewRegistry().Add(tt.rule); err == nil {
				t.Errorf("Add(%+v) returned no error", tt.rule)
			}
		})
	}
	if err := NewRegistry().Add(nil); err == nil {
		t.Error("Add(nil) returned no error")
	}
	var registry *Registry
	if err := registry.Add(Rule{ID: "id", Description: "rule", Pattern: regexp.MustCompile("x")}); err == nil {
		t.Error("Add on nil registry returned no error")
	}
}

func TestRegistryMatchesInRegistrationOrder(t *testing.T) {
	registry := NewRegistry()
	for _, rule := range []Rule{
		{ID: "first", Description: "first match", Pattern: regexp.MustCompile("secret")},
		{ID: "miss", Description: "does not match", Pattern: regexp.MustCompile("other")},
		{ID: "second", Description: "second match", Pattern: regexp.MustCompile("secret")},
	} {
		if err := registry.Add(rule); err != nil {
			t.Fatal(err)
		}
	}
	got := registry.Matches("a secret here")
	want := []string{"first match", "second match"}
	if len(got) != len(want) {
		t.Fatalf("Matches() = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Matches()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestRuleMatchHonorsEntropyAndNilPattern(t *testing.T) {
	if (Rule{}).Match("anything") {
		t.Error("rule with nil pattern matched text")
	}
	pattern := regexp.MustCompile("aaaa")
	if (Rule{Pattern: pattern, MinEntropy: 1}).Match("aaaa") {
		t.Error("low-entropy match passed the threshold")
	}
	if !(Rule{Pattern: pattern}).Match("aaaa") {
		t.Error("pattern-only rule failed to match")
	}
}

func TestNilRegistryMatchesNothing(t *testing.T) {
	var registry *Registry
	if got := registry.Matches("secret"); got != nil {
		t.Errorf("nil registry Matches() = %v, want nil", got)
	}
}
