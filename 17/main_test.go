package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	input := []byte(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)

	registers, opcodes := parseInput(string(input))
	assert.Equal(t, []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}, partOne(registers, opcodes))
}

func TestPart1_1(t *testing.T) {
	registers := []int{0, 0, 9}
	opcodes := []int{2, 6}
	partOne(registers, opcodes)
	assert.Equal(t, 1, registers[1])
}

func TestPart1_2(t *testing.T) {
	registers := []int{10, 0, 0}
	opcodes := []int{5, 0, 5, 1, 5, 4}
	output := partOne(registers, opcodes)
	assert.Equal(t, []int{0, 1, 2}, output)
}

func TestPart1_4(t *testing.T) {
	registers := []int{0, 29, 0}
	opcodes := []int{1, 7}
	partOne(registers, opcodes)
	assert.Equal(t, 26, registers[1])
}

func TestPart1_5(t *testing.T) {
	registers := []int{0, 2024, 43690}
	opcodes := []int{4, 0}
	partOne(registers, opcodes)
	assert.Equal(t, 44354, registers[1])
}
