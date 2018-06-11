package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestByVersion(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "nothing",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "invalid",
			input: []string{"x", "y", "z"},
			want:  []string{"x", "y", "z"},
		},
		{
			name:  "major versions",
			input: []string{"v3.0", "v1.0", "v2.0"},
			want:  []string{"v3.0", "v2.0", "v1.0"},
		},
		{
			name:  "minor versions",
			input: []string{"v0.1", "v0.3", "v0.2"},
			want:  []string{"v0.3", "v0.2", "v0.1"},
		},
		{
			name:  "minor versions padding",
			input: []string{"v0.1.2", "v0.1.1", "v0.1"},
			want:  []string{"v0.1.2", "v0.1.1", "v0.1"},
		},
		{
			name:  "minor versions numeric",
			input: []string{"v1.0.19", "v1.0.2"},
			want:  []string{"v1.0.19", "v1.0.2"},
		},
		{
			name:  "multiple padding",
			input: []string{"v1.0", "v2.2.2", "v2.1", "v0.0.0.999", "v3.5"},
			want:  []string{"v3.5", "v2.2.2", "v2.1", "v1.0", "v0.0.0.999"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rel := make([]Release, 0, len(test.input))
			for _, tag := range test.input {
				rel = append(rel, Release{TagName: tag})
			}
			sort.Sort(byVersion(rel))
			got := make([]string, 0, len(test.input))
			for _, r := range rel {
				got = append(got, r.TagName)
			}

			if len(got) != len(test.want) || !reflect.DeepEqual(got, test.want) {
				t.Errorf("byVersion(%v): got: %v, want: %v", test.input, got, test.want)
			}
		})
	}
}
