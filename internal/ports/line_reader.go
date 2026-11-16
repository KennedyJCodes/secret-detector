package ports

import "github.com/KennedyJCodes/secret-detector/internal/scanner"

// LineReader reads a source line by line and visits each line.
type LineReader interface {
	ReadLines(source Source, visit func(scanner.Line) error) error
}

// Source abstracts a byte source such as a file or an in-memory buffer.
type Source interface {
	Read(p []byte) (n int, err error)
}
