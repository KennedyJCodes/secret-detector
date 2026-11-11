// Package bufioreader adapts bufio.Scanner to the scanner.LineReader port.
package bufioreader

import (
	"bufio"
	"fmt"

	"github.com/KennedyJCodes/secret-detector/internal/scanner"
)

// Config controls the maximum line size. MaxLineBytes is required and must be
// greater than zero. The limit avoids relying on bufio.Scanner's 64 KiB default.
type Config struct {
	MaxLineBytes int
}

type Reader struct {
	maxLineBytes int
}

func New(config Config) (*Reader, error) {
	if config.MaxLineBytes <= 0 {
		return nil, fmt.Errorf("MaxLineBytes must be greater than zero")
	}
	return &Reader{maxLineBytes: config.MaxLineBytes}, nil
}

func (r *Reader) ReadLines(source scanner.Source, visit func(scanner.Line) error) error {
	if source == nil {
		return fmt.Errorf("source cannot be nil")
	}
	if visit == nil {
		return fmt.Errorf("visit function cannot be nil")
	}

	s := bufio.NewScanner(source)
	// The initial buffer must not exceed the configured maximum: Scanner may
	// accept tokens that fit in the buffer even if they exceed max.
	initialBufferSize := r.maxLineBytes
	if initialBufferSize > 64*1024 {
		initialBufferSize = 64 * 1024
	}
	s.Buffer(make([]byte, initialBufferSize), r.maxLineBytes)
	lineNumber := 0
	for s.Scan() {
		lineNumber++
		if err := visit(scanner.Line{Number: lineNumber, Text: s.Text()}); err != nil {
			return fmt.Errorf("processing line %d: %w", lineNumber, err)
		}
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("reading lines: %w", err)
	}
	return nil
}
