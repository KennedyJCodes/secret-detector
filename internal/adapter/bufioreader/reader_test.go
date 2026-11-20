package bufioreader

import (
	"errors"
	"strings"
	"testing"

	"github.com/KennedyJCodes/secret-detector/internal/scanner"
)

func TestNewRequiresPositiveMaximumLineSize(t *testing.T) {
	for _, size := range []int{0, -1} {
		if reader, err := New(Config{MaxLineBytes: size}); err == nil || reader != nil {
			t.Errorf("New(MaxLineBytes=%d) = (%v, %v), want nil reader and error", size, reader, err)
		}
	}
	if _, err := New(Config{MaxLineBytes: 1}); err != nil {
		t.Fatalf("New(MaxLineBytes=1) returned error: %v", err)
	}
}

func TestReadLinesVisitsLinesInOrder(t *testing.T) {
	reader, err := New(Config{MaxLineBytes: 64})
	if err != nil {
		t.Fatal(err)
	}
	var got []scanner.Line
	err = reader.ReadLines(strings.NewReader("first\nsecond\n"), func(line scanner.Line) error {
		got = append(got, line)
		return nil
	})
	if err != nil {
		t.Fatalf("ReadLines returned error: %v", err)
	}
	want := []scanner.Line{{Number: 1, Text: "first"}, {Number: 2, Text: "second"}}
	if len(got) != len(want) {
		t.Fatalf("ReadLines visited %d lines, want %d: %#v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("line %d = %#v, want %#v", i, got[i], want[i])
		}
	}
}

func TestReadLinesRejectsNilInputs(t *testing.T) {
	reader, err := New(Config{MaxLineBytes: 64})
	if err != nil {
		t.Fatal(err)
	}
	if err := reader.ReadLines(nil, func(scanner.Line) error { return nil }); err == nil {
		t.Error("ReadLines with nil source returned no error")
	}
	if err := reader.ReadLines(strings.NewReader("line"), nil); err == nil {
		t.Error("ReadLines with nil visitor returned no error")
	}
}

func TestReadLinesWrapsVisitorErrorAndStops(t *testing.T) {
	reader, err := New(Config{MaxLineBytes: 64})
	if err != nil {
		t.Fatal(err)
	}
	wantErr := errors.New("stop")
	visited := 0
	err = reader.ReadLines(strings.NewReader("first\nsecond\n"), func(scanner.Line) error {
		visited++
		return wantErr
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("ReadLines error = %v, want wrapped %v", err, wantErr)
	}
	if visited != 1 {
		t.Errorf("visitor called %d times, want 1", visited)
	}
}

func TestReadLinesRejectsLineAboveMaximum(t *testing.T) {
	reader, err := New(Config{MaxLineBytes: 4})
	if err != nil {
		t.Fatal(err)
	}
	err = reader.ReadLines(strings.NewReader("12345\n"), func(scanner.Line) error { return nil })
	if err == nil || !strings.Contains(err.Error(), "reading lines") {
		t.Fatalf("ReadLines error = %v, want scanner line-size error", err)
	}
}
