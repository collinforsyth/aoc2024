package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

func TestPart1Sample(t *testing.T) {
	input := []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)
	parsed := parseInput(util.FromBytes(input))
	assert.Equal(t, 143, partOne(parsed))
}

func BenchmarkPart1(b *testing.B) {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		b.Error(err)
	}
	r := 0
	parsed := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = partOne(parsed)
	}
	result = r
}

func TestPart2Sample(t *testing.T) {
	input := []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)
	parsed := parseInput(util.FromBytes(input))
	assert.Equal(t, 123, partTwo(parsed))
}

func BenchmarkPart2(b *testing.B) {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		b.Error(err)
	}
	r := 0
	parsed := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = partTwo(parsed)
	}
	result = r
}
