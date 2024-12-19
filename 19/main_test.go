package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	input := []byte(`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`)
	dict, words := parseInput(util.FromBytes(input).Lines())
	assert.Equal(t, 6, partOne(dict, words))
}

func TestPart2Sample(t *testing.T) {
	input := []byte(`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`)
	dict, words := parseInput(util.FromBytes(input).Lines())
	assert.Equal(t, 16, partTwo(dict, words))
}
