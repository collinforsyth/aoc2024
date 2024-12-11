package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	input := "125 17"
	parsed := parseInput(input)
	assert.Equal(t, 22, solve(parsed, 6))
	parsed = parseInput(input)
	assert.Equal(t, 55312, solve(parsed, 25))
}

func BenchmarkPart2(b *testing.B) {
	input, err := util.ReadInput("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	parsed := parseInput(input.String())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(parsed, 75)
	}

}
