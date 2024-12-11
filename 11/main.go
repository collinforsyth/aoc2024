package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
	"math"
	"slices"
	"strings"
)

func main() {
	input, err := util.ReadInput("11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parsed := parseInput(input.String())
	fmt.Println("Part 1: ", partOne(parsed, 25))
}

func parseInput(input string) []int {
	a := make([]int, 0)
	for _, s := range strings.Fields(input) {
		a = append(a, util.MustAtoi(s))
	}
	return a
}

func partOne(input []int, blinks int) int {
	for range blinks {
		for i := 0; i < len(input); i++ {
			if input[i] == 0 {
				input[i] = 1
				continue
			}
			if d := digits(input[i]); d%2 == 0 {
				a, b := split(input[i], d)
				input = slices.Replace(input, i, i+1, a, b)
				i++
				continue
			}
			input[i] = input[i] * 2024
		}
	}
	return len(input)
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
