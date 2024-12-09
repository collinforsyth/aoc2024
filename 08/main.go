package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
)

func main() {
	input, err := util.ReadInput("08/input.txt")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Part 1: ", partOne(input.Runes()))
	fmt.Println("Part 2: ", partTwo(input.Runes()))
}

func partOne(input [][]rune) int {
	m := make(map[rune][]point, 0)
	for i := range input {
		for j := range input[i] {
			if input[i][j] != '.' {
				m[input[i][j]] = append(m[input[i][j]], point{i, j})
			}
		}
	}

	sum := make(map[point]struct{})
	for _, antennas := range m {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				d := distance(antennas[i], antennas[j])
				a1, a2 := antinodes(antennas[i], antennas[j], d)
				if check(input, a1.x, a1.y) {
					sum[a1] = struct{}{}
				}
				if check(input, a2.x, a2.y) {
					sum[a2] = struct{}{}
				}

			}
		}
	}
	return len(sum)
}

func partTwo(input [][]rune) int {
	m := make(map[rune][]point, 0)
	for i := range input {
		for j := range input[i] {
			if input[i][j] != '.' {
				m[input[i][j]] = append(m[input[i][j]], point{i, j})
			}
		}
	}

	sum := make(map[point]struct{})
	for _, antennas := range m {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				a1, a2 := antennas[i], antennas[j]
				sum[a1], sum[a2] = struct{}{}, struct{}{}
				d := distance(a1, a2)
				// continue iterating down the line until we hit walls
				// on both sides.
				c1, c2 := true, true
				for c1 || c2 {
					a1, a2 = antinodes(a1, a2, d)
					c1, c2 = check(input, a1.x, a1.y), check(input, a2.x, a2.y)
					if c1 {
						sum[a1] = struct{}{}
					}
					if c2 {
						sum[a2] = struct{}{}
					}
				}
			}
		}
	}
	return len(sum)
}

type point struct {
	x, y int
}

type vector point

func antinodes(a, b point, d vector) (point, point) {
	// check diagonal position
	if (d.x < 0 && d.y < 0) || (d.x > 0 && d.y > 0) {
		a1 := point{a.x + d.x, a.y + d.y}
		a2 := point{b.x - d.x, b.y - d.y}
		return a1, a2
	}
	a1 := point{a.x + d.x, a.y + d.y}
	a2 := point{b.x - d.x, b.y - d.y}
	return a1, a2
}

func distance(a, b point) vector {
	return vector{a.x - b.x, a.y - b.y}
}

func check(input [][]rune, i, j int) bool {
	return i >= 0 && i < len(input) && j >= 0 && j < len(input[i])
}
