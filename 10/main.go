package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
)

func main() {
	input, err := util.ReadInput("10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid := parseInput(input.Runes())
	fmt.Println("Part 1: ", partOne(grid))
	fmt.Println("Part 2: ", partTwo(grid))
}

func parseInput(input [][]rune) [][]int {
	a := make([][]int, len(input))
	for i := range input {
		a[i] = make([]int, len(input[i]))
		for j := range input[i] {
			a[i][j] = int(input[i][j] - '0')
		}
	}
	return a
}

func partOne(grid [][]int) int {
	s := 0
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 0 {
				s += bfs(grid, point{x: j, y: i}, true)
			}
		}
	}
	return s
}

func partTwo(grid [][]int) int {
	s := 0
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 0 {
				s += bfs(grid, point{x: j, y: i}, false)
			}
		}
	}
	return s
}

type point struct {
	x, y int
}

var directions = []point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func bfs(grid [][]int, p point, markSeen bool) int {
	seen := make(map[point]bool)
	sum := 0
	q := util.NewQueue[point]()
	seen[p] = true
	q.Push(p)
	for q.Len() > 0 {
		p, _ := q.Pop()
		if grid[p.y][p.x] == 9 {
			sum++
		}
		for _, d := range directions {
			p1 := point{x: p.x + d.x, y: p.y + d.y}
			if !check(grid, p1) {
				continue
			}
			if markSeen && seen[p1] {
				continue
			}
			if grid[p1.y][p1.x] == grid[p.y][p.x]+1 {
				q.Push(p1)
				seen[p1] = true
			}
		}
	}
	return sum
}

func check(grid [][]int, p point) bool {
	return p.x >= 0 && p.x < len(grid) && p.y >= 0 && p.y < len(grid[0])
}
