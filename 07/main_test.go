package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	result int
)

func Test_concat(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 12},
		{20, 1, 201},
	}

	for i, tc := range testCases {
		assert.Equal(t, tc.expected, concat(tc.a, tc.b), "testCases[%d] - failed", i)
	}
}

func TestPart1Sample(t *testing.T) {
	input := []byte(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

	equations := parseInput(util.FromBytes(input))
	assert.Equal(t, 3749, partOne(equations))
}

func TestPart2Sample(t *testing.T) {
	input := []byte(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

	equations := parseInput(util.FromBytes(input))
	assert.Equal(t, 11387, partTwo(equations))
}

func BenchmarkPart1(b *testing.B) {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		b.Error(err)
	}
	r := 0
	equations := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = partOne(equations)
	}
	result = r
}

func BenchmarkPart2(b *testing.B) {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		b.Error(err)
	}
	r := 0
	equations := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = partTwo(equations)
	}
	result = r
}
