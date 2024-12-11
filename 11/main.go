package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := util.ReadInput("11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parsed := parseInput(input.String())
	fmt.Println("Part 1: ", solve(parsed, 25))
	parsed = parseInput(input.String())
	fmt.Println("Part 2: ", solve(parsed, 75))
}

func parseInput(input string) map[int]int {
	m := make(map[int]int, 0)
	for _, s := range strings.Fields(input) {
		m[util.MustAtoi(s)]++
	}
	return m
}

func solve(input map[int]int, blinks int) int {
	for range blinks {
		curr := make(map[int]int, len(input))
		for v, c := range input {
			if v == 0 {
				curr[1] += c
			} else if digits(v)%2 == 0 {
				a, b := split(v, digits(v))
				curr[a] += c
				curr[b] += c
			} else {
				curr[v*2024] += c
			}
		}
		input = curr
	}
	sum := 0
	for _, c := range input {
		sum += c
	}
	return sum
}

// https://stackoverflow.com/a/68124773
func digits(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

func split(i int, digits int) (int, int) {
	m := digits / 2
	p := int(math.Pow10(m))
	return i / p, i % p
}
