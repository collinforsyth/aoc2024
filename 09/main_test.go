package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

func TestPart1Sample(t *testing.T) {
	input := "12345"
	partOne(parseInput(input))

	input2 := "2333133121414131402"
	assert.Equal(t, 1928, partOne(parseInput(input2)))
}

func BenchmarkPart1(b *testing.B) {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		in := parseInput(input.String())
		b.StartTimer()
		result = partOne(in)
	}
}

func TestPart2Sample(t *testing.T) {
	input := "2333133121414131402"
	assert.Equal(t, 2858, partTwo(parseInput(input)))
}

func BenchmarkPart2(b *testing.B) {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		in := parseInput(input.String())
		b.StartTimer()
		result = partTwo(in)
	}
}
