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
		count     int
		input     []byte
		direction point
	}{
		{
			name:  "right",
			count: 1,
			input: []byte(`....
XMAS
....
....`),
		},
		{
			name:  "left",
			count: 1,
			input: []byte(`....
SAMX
....
....`),
		},
		{
			name:  "down",
			count: 1,
			input: []byte(`.X..
.M..
.A..
.S..`),
		},
		{
			name:  "up",
			count: 1,
			input: []byte(`.S..
.A..
.M..
.X..`),
		},
		{
			name:  "diagonal/down-right",
			count: 1,
			input: []byte(`X...
.M..
..A.
...S`),
		},
		{
			name:  "diagonal/down-left",
			count: 1,
			input: []byte(`...X
..M.
.A..
S...`),
		},
		{
			name:  "diagonal/up-right",
			count: 1,
			input: []byte(`...S
..A.
.M..
X...`),
		},
		{
			name:  "diagonal/up-left",
			count: 1,
			input: []byte(`S...
.A..
..M.
...X`),
		},
		{
			name:  "multiple",
			count: 8,
			input: []byte(`S..S..S
.A.A.A.
..MMM..
SAMXMAS
..MMM..
.A.A.A.
S..S..S`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsed := parseInput(tc.input)
			assert.Equal(t, tc.count, partOne(parsed))
		})
	}
}
