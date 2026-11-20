package app

import (
	"fmt"
	"io"
	"strings"

	"github.com/KennedyJCodes/secret-detector/internal/ports"
	"github.com/KennedyJCodes/secret-detector/internal/scanner"
)

// Run scans the source and writes only matching lines with their one-based line number.
func Run(lineReader ports.LineReader, source ports.Source, output io.Writer, ruleSet ports.RuleSet) error {
	if output == nil {
		return fmt.Errorf("output cannot be nil")
	}
	if lineReader == nil {
		return fmt.Errorf("line reader cannot be nil")
	}
	if ruleSet == nil {
		return fmt.Errorf("rule set cannot be nil")
	}
	return lineReader.ReadLines(source, func(line scanner.Line) error {
		matches := ruleSet.Matches(line.Text)
		if len(matches) == 0 {
			return nil
		}
		_, err := fmt.Fprintf(output, "%d: [%s detected] %s\n", line.Number, strings.Join(matches, ", "), line.Text)
		return err
	})
}
