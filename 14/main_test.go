package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []byte(`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`)
	robots := parseInput(util.FromBytes(input).Lines())
	assert.Equal(t, 12, partOne(11, 7, robots))
}

func TestTravel(t *testing.T) {
	expected := []point{
		{x: 2, y: 4},
		{x: 4, y: 1},
		{x: 6, y: 5},
		{x: 8, y: 2},
		{x: 10, y: 6},
		{x: 1, y: 3},
	}

	x, y := 11, 7
	v := point{x: 2, y: -3}

	for i := 1; i < len(expected); i++ {
		assert.Equal(t, expected[i], next(x, y, expected[i-1], v))
	}
}
