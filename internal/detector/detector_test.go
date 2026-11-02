package detector_test

import (
	"slices"
	"testing"

	"github.com/ajpotts01/secret-detector/internal/config"
	"github.com/ajpotts01/secret-detector/internal/detector"
)

func TestRedaction(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []detector.Finding
	}{
		{
			name: "Plain line",
			line: "Howdy!",
			want: nil,
		},
		{
			name: "AWS key",
			line: "AKIAIOSFODNN7EXAMPLE", //secret-detector:ignore
			want: []detector.Finding{
				{
					RuleID:         "aws-key",
					FilePath:       "test.txt",
					LineNumber:     1,
					RedactedSecret: "AKIA...MPLE",
				},
			},
		},
		{
			name: "Ignored line",
			line: "key := AKIAIOSFODNN7EXAMPLE //secret-detector:ignore",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := detector.NewEngine(config.DefaultRules())
			got := e.InspectLine("test.txt", 1, tt.line)

			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
