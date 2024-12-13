package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	input := []byte(`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`)

	parsed := parseInput(util.FromBytes(input).Lines(), 0)
	assert.Equal(t, 480, partOne(parsed))

	parsed = parseInput(util.FromBytes(input).Lines(), 10000000000000)
	t.Log(partOne(parsed))
}
