package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	input := "12345"
	partOne(parseInput(input))

	input2 := "2333133121414131402"
	assert.Equal(t, 1928, partOne(parseInput(input2)))
}

func TestPart2Sample(t *testing.T) {
	input := "2333133121414131402"
	assert.Equal(t, 2858, partTwo(parseInput(input)))
}
