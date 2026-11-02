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
			Data: []byte("Hello!\nAKIAIOSFODNN7EXAMPLE"), //secret-detector:ignore
		},
	}

	var got []string

	err := input.ReadSingleFile(fs, "test.txt", func(filePath string, lineNum int, line string) {
		got = append(got, line)
	})

	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	want := []string{"Hello!", "AKIAIOSFODNN7EXAMPLE"} //secret-detector:ignore
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReadInputDir(t *testing.T) {
	tests := []struct {
		name string
		fs   fstest.MapFS
		want []string
	}{
		{
			name: "reads files in nested directories",
			fs: fstest.MapFS{
				"a.txt": &fstest.MapFile{
					Data: []byte("Hello from file a.txt!"),
				},
				"sub/b.txt": &fstest.MapFile{
					Data: []byte("Hello from file sub/b.txt!"),
				},
			},
			want: []string{"Hello from file a.txt!", "Hello from file sub/b.txt!"},
		},
		{
			name: "skips ignored directories like .git",
			fs: fstest.MapFS{
				"c.txt": &fstest.MapFile{
					Data: []byte("Hello from file c.txt!"),
				},
				".git/config": &fstest.MapFile{
					Data: []byte(".git/config should be skipped."),
				},
			},
			want: []string{"Hello from file c.txt!"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			err := input.ReadDir(tt.fs, ".", func(filePath string, lineNum int, line string) {
				got = append(got, line)
			})

			if err != nil {
				t.Fatalf("error reading files: %v", err)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
