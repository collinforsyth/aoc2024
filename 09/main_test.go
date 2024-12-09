package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	input := "12345"
	partOne(input)

	input2 := "2333133121414131402"
	assert.Equal(t, 1928, partOne(input2))
}
