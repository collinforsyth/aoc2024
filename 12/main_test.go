package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []byte(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)

	assert.Equal(t, 1930, partOne(util.FromBytes(input).Runes()))
}

func TestPart2(t *testing.T) {
	t.Skip()
	input := []byte(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`)
	assert.Equal(t, 368, partOne(util.FromBytes(input).Runes()))
}
