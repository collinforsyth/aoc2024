package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestPart1Sample(t *testing.T) {
	input := []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	parsed := parseInput(input)
	assert.Equal(t, 18, partOne(parsed))
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		name      string
		input     []byte
		direction point
	}{
		{
			name: "right",
			input: []byte(`....
XMAS
....
....`),
		},
		{
			name: "left",
			input: []byte(`....
SAMX
....
....`),
		},
		{
			name: "down",
			input: []byte(`.X..
.M..
.A..
.S..`),
		},
		{
			name: "up",
			input: []byte(`.S..
.A..
.M..
.X..`),
		},
		{
			name: "diagonal/down-right",
			input: []byte(`X...
.M..
..A.
...S`),
		},
		{
			name: "diagonal/down-left",
			input: []byte(`...X
..M.
.A..
S...`),
		},
		{
			name: "diagonal/up-right",
			input: []byte(`...S
..A.
.M..
X...`),
		},
		{
			name: "diagonal/up-left",
			input: []byte(`S...
.A..
..M.
...X`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsed := parseInput(tc.input)
			assert.Equal(t, 1, partOne(parsed))
		})
	}
}
