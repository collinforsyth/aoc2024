package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	input := []byte(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)
	parsed := util.FromBytes(input).Runes()
	assert.Equal(t, 41, partOne(parsed))
}

func TestPart2Sample(t *testing.T) {
	input := []byte(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)
	parsed := util.FromBytes(input).Runes()
	assert.Equal(t, 6, partTwo(parsed))
}