package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
)

func main() {
	input, err := util.ReadInput("06/input.txt")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Part 1: ", partOne(input.Clone().Runes()))
	fmt.Println("Part 2: ", partTwo(input.Clone().Runes()))

}

func partOne(input [][]rune) int {
	sum := 0
	i, j := getStart(input)
	d := point{0, -1}
	for check(input, i, j) {
		if input[i][j] == '.' {
			input[i][j] = 'X'
			sum++
		}
		if input[i][j] == '#' {
			// back up one step
			i -= d.y
			j -= d.x
			d = rotate(d)
		}
		i += d.y
		j += d.x
	}
	// mark one more 'X' on the way out
	return sum + 1
}

type vector struct {
	x, y, dx, dy int
}

func partTwo(input [][]rune) int {
	sum := 0
	i, j := getStart(input)
	d := point{0, -1}
	vectors := make(map[vector]struct{})
	vectors[vector{x: j, y: i, dx: d.x, dy: d.y}] = struct{}{}
	for check(input, i, j) {
		// check if we've visited this point before where
		next := rotate(d)
		if _, ok := vectors[vector{x: j, y: i, dx: next.x, dy: next.y}]; ok {
			sum++
		}
		vectors[vector{x: i, y: j, dx: d.x, dy: d.y}] = struct{}{}
		if input[i][j] == '#' {
			// back up one step
			i -= d.y
			j -= d.x
			d = rotate(d)
		}
		i += d.y
		j += d.x
	}
	return sum
}

func getStart(input [][]rune) (int, int) {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == '^' {
				return i, j
			}
		}
	}
	panic("not found")
}

type point struct {
	x, y int
}

func rotate(d point) point {
	return point{x: -d.y, y: d.x}
}

func check(input [][]rune, i, j int) bool {
	if i < 0 || j < 0 || i >= len(input) || j >= len(input[i]) {
		return false
	}
	return true
}
