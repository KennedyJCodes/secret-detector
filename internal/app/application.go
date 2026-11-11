package app

import (
	"fmt"
	"io"

	"github.com/KennedyJCodes/secret-detector/internal/scanner"
)

// Run scans the source and writes each line with its one-based line number.
func Run(lineReader scanner.LineReader, source scanner.Source, output io.Writer) error {
	if output == nil {
		return fmt.Errorf("output cannot be nil")
	}
	return lineReader.ReadLines(source, func(line scanner.Line) error {
		_, err := fmt.Fprintf(output, "%d: %s\n", line.Number, line.Text)
		return err
	})
}
