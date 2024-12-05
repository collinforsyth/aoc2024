package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
	"slices"
	"strings"
)

type instructions struct {
	rules   map[int][]int
	updates [][]int
}

func main() {
	input, err := util.ReadInput("05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parsed := parseInput(input)
	fmt.Println("Part 1: ", partOne(parsed))
	fmt.Println("Part 2: ", partTwo(parsed))
}

func parseInput(input *util.Input) instructions {
	parsed := instructions{
		rules:   make(map[int][]int),
		updates: make([][]int, 0),
	}
	for line := range input.Lines() {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == '|'
		})
		switch len(fields) {
		case 0:
			continue
		case 2:
			t1, t2 := util.MustAtoi(fields[0]), util.MustAtoi(fields[1])
			parsed.rules[t1] = append(parsed.rules[t1], t2)
		default:
			t := make([]int, 0)
			for _, s := range fields {
				t = append(t, util.MustAtoi(s))
			}
			parsed.updates = append(parsed.updates, t)
		}
	}
	return parsed
}

func partOne(instr instructions) int {
	sum := 0
	for _, update := range instr.updates {
		valid := true
		for i := 0; i < len(update); i++ {
			for j := i; j < len(update); j++ {
				if incorrectOrder(instr.rules, update[j], update[i]) {
					valid = false
				}
			}
		}
		if valid {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func partTwo(instr instructions) int {
	sum := 0
	for _, update := range instr.updates {
		fixed := false
		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				if incorrectOrder(instr.rules, update[j], update[i]) {
					// swap
					update[j], update[i] = update[i], update[j]
					fixed = true
				}
			}
		}
		if fixed {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func incorrectOrder(rules map[int][]int, a, b int) bool {
	return slices.Contains(rules[a], b)
}
