package scanner

// Line represents a line read from a text source. Number starts at 1.
type Line struct {
	Number int
	Text   string
}

// LineReader is the input port used by the scanning application.
// Its implementation can change without coupling the application to bufio.
type LineReader interface {
	ReadLines(source Source, visit func(Line) error) error
}

// Source abstracts the origin of bytes, such as a file or an in-memory buffer.
// It matches io.Reader, Go's standard input port for byte sources.
type Source interface {
	Read(p []byte) (n int, err error)
}
