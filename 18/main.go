package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"iter"
	"strings"
)

func main() {
	input, err := util.ReadInput("18/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	falling := parseInput(input.Lines())
	grid := newGrid(71, 71, '.')
	fmt.Println("Part 1: ", partOne(grid, falling, 1024, point{x: 70, y: 70}))
	fmt.Println("Part 2: ", partTwo(71, 71, falling, point{x: 70, y: 70}))
}

type point struct {
	x, y int
}

type search struct {
	p    point
	dist int
}

var directions = []point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func partTwo(x, y int, falling []point, toFind point) string {
	start, end := 0, len(falling)
	for start <= end {
		mid := (start + end) / 2
		res := partOne(newGrid(x, y, '.'), falling, mid, toFind)
		if res == -1 {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	r := max(start, end)
	return fmt.Sprintf("%d,%d", falling[r-1].x, falling[r-1].y)
}

func partOne(grid [][]rune, falling []point, marked int, toFind point) int {
	for i := 0; i < marked; i++ {
		mark(grid, falling[i], '#')
	}
	seen := make(map[point]bool)
	p := point{x: 0, y: 0}
	seen[p] = true
	q := util.NewQueue[search]()
	q.Push(search{p: p, dist: 0})
	for q.Len() > 0 {
		s, _ := q.Pop()
		for _, d := range directions {
			p1 := point{x: s.p.x + d.x, y: s.p.y + d.y}
			if !check(grid, p1) || grid[p1.y][p1.x] == '#' || seen[p1] {
				continue
			}
			if p1 == toFind {
				return s.dist + 1
			}
			q.Push(search{
				p:    p1,
				dist: s.dist + 1,
			})
			seen[p1] = true
		}
	}
	return -1
}

func parseInput(input iter.Seq2[int, string]) []point {
	a := make([]point, 0)
	for _, s := range input {
		split := strings.FieldsFunc(s, func(r rune) bool {
			return r == ','
		})
		a = append(a, point{
			x: util.MustAtoi(split[0]),
			y: util.MustAtoi(split[1]),
		})
	}
	return a
}

func newGrid(x, y int, v rune) [][]rune {
	grid := make([][]rune, y)
	for i := range grid {
		grid[i] = make([]rune, x)
		for j := range grid[i] {
			grid[i][j] = v
		}
	}
	return grid
}

func mark(grid [][]rune, p point, c rune) {
	grid[p.y][p.x] = c
}

func check(grid [][]rune, p point) bool {
	return p.x >= 0 && p.x < len(grid[0]) && p.y >= 0 && p.y < len(grid)
}
