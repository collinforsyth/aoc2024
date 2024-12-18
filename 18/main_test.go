package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	input := []byte(`5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`)

	falling := parseInput(util.FromBytes(input).Lines())
	grid := newGrid(7, 7, '.')
	assert.Equal(t, 22, partOne(grid, falling, 12, point{x: 6, y: 6}))
}
