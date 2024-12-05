package main

import (
	"collinforsyth/aoc2024/util"
	"log"
	"testing"

	"gotest.tools/assert"
)

var result int

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

func BenchmarkPart1(b *testing.B) {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		log.Println(err)
		return
	}
	r := 0
	parsed := parseInput(input.Bytes())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = partOne(parsed)
	}
	result = r
}

func TestPart2Sample(t *testing.T) {
	input := []byte(`.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`)
	parsed := parseInput(input)
	assert.Equal(t, 9, partTwo(parsed))
}

func BenchmarkPart2(b *testing.B) {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		log.Println(err)
		return
	}
	parsed := parseInput(input.Bytes())
	r := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = partTwo(parsed)
	}
	result = r
}
