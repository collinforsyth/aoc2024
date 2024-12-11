package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	input := []byte(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)

	grid := parseInput(util.FromBytes(input).Runes())
	assert.Equal(t, 36, partOne(grid))
}

func TestPart2Sample(t *testing.T) {
	input := []byte(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)

	grid := parseInput(util.FromBytes(input).Runes())
	assert.Equal(t, 81, partTwo(grid))
}
