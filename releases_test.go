package main

import (
	"reflect"
	"testing"
)

func TestTagVersion(t *testing.T) {
	tests := []struct {
		name    string
		tag     string
		wantErr bool
		want    []int
	}{
		{
			name:    "no_v",
			tag:     "1.2.3",
			wantErr: true,
		},
		{
			name:    "only_v",
			tag:     "v",
			wantErr: true,
		},
		{
			name:    "trailing_dot",
			tag:     "v1.2.3.",
			wantErr: true,
		},
		{
			name:    "multiple_dot",
			tag:     "v1.2..3",
			wantErr: true,
		},
		{
			name:    "non_numeric",
			tag:     "v1.2.x",
			wantErr: true,
		},
		{
			name: "0.15",
			tag:  "v0.15",
			want: []int{0, 15},
		},
		{
			name: "1.2.3",
			tag:  "v1.2.3",
			want: []int{1, 2, 3},
		},
		{
			name: "1.0.12.987",
			tag:  "v1.0.12.987",
			want: []int{1, 0, 12, 987},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Release{TagName: test.tag}
			got, err := r.TagVersion()

			if test.wantErr {
				if err == nil {
					t.Errorf("TagVersion(%v): got: %v, nil, want: nil, err", test.tag, got)
				}
			} else {
				if err != nil || !reflect.DeepEqual(got, test.want) {
					t.Errorf("TagVersion(%v): got: %v, %v, want: %v, nil", test.tag, got, err, test.want)
				}
			}
		})
	}
}
