package app

import (
	"fmt"
	"io"
	"strings"

	"github.com/KennedyJCodes/secret-detector/internal/ports"
	"github.com/KennedyJCodes/secret-detector/internal/scanner"
)

// Run scans the source and writes each line with its one-based line number.
func Run(lineReader ports.LineReader, source ports.Source, output io.Writer, rules ...ports.Rule) error {
	if output == nil {
		return fmt.Errorf("output cannot be nil")
	}
	if lineReader == nil {
		return fmt.Errorf("line reader cannot be nil")
	}
	for _, rule := range rules {
		if rule == nil {
			return fmt.Errorf("rule cannot be nil")
		}
	}
	return lineReader.ReadLines(source, func(line scanner.Line) error {
		var matches []string
		for _, rule := range rules {
			if rule.Match(line.Text) {
				matches = append(matches, rule.Name())
			}
		}
		if len(matches) == 0 {
			return nil
		}
		_, err := fmt.Fprintf(output, "%d: [%s detected] %s\n", line.Number, strings.Join(matches, ", "), line.Text)
		return err
	})
}
