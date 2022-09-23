package entities

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *Log
		wantErr  bool
	}{
		{
			name:  "test",
			input: "{\"file\":'scf20220910.html.gz',\"size\":10965}",
			expected: &Log{
				File: "scf20220910.html.gz",
				Size: 10965,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unmarshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.expected); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
