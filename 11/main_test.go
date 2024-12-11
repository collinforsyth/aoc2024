package main

import (
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
