package util

import (
	"testing"

	"gotest.tools/assert"
)

func TestInput(t *testing.T) {
	testCases := []struct {
		name  string
		bytes string
		want  []string
	}{
		{
			name: "simple",
			bytes: `line 1
line 2
line 3`,
			want: []string{"line 1", "line 2", "line 3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := Input{input: tc.bytes}
			i := 0
			for v := range input.Lines() {
				assert.Equal(t, tc.want[i], v)
				i++
			}
		})
	}
}
