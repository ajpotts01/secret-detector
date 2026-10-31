package input_test

import (
	"slices"
	"testing"
	"testing/fstest"

	"github.com/ajpotts01/secret-detector/internal/input"
)

func TestReadInputSingle(t *testing.T) {
	fs := fstest.MapFS{
		"test.txt": &fstest.MapFile{
			Data: []byte("Hello!\nAKIAIOSFODNN7EXAMPLE"),
		},
	}

	var got []string

	err := input.ReadInputSingle(fs, "test.txt", func(lineNum int, line string) {
		got = append(got, line)
	})

	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	want := []string{"Hello!", "AKIAIOSFODNN7EXAMPLE"}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}