package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
)

func main() {
	input, err := util.ReadInput("12/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Part 1:", partOne(input.Runes()))
}

func partOne(grid [][]rune) int {
	price := 0
	seen := make(map[point]bool)
	for i := range grid {
		for j := range grid[i] {
			p := point{x: j, y: i}
			if seen[p] {
				continue
			}
			price += floodFill(grid, seen, point{x: j, y: i})
		}
	}
	return price
}

type point struct {
	x, y int
}

var directions = []point{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

// +X+-+-+-+
// XAXA A A|
// +X+-+-+-+
type plot struct {
	perimeter int
	area      int
	vertices  int
}

func floodFill(grid [][]rune, seen map[point]bool, p point) int {
	var (
		plt   plot
		plant = grid[p.y][p.x]
		q     = util.NewQueue[point]()
	)
	seen[p] = true
	q.Push(p)
	for q.Len() > 0 {
		p, _ = q.Pop()
		plt.area++
		for _, d := range directions {
			p1 := point{x: p.x + d.x, y: p.y + d.y}
			if !check(grid, p1) {
				// anything on the outside of the grid
				// adds to the perimeter
				plt.perimeter++
				continue
			}
			v := grid[p1.y][p1.x]
			if v != plant {
				// if it's not the same plant, it's part of the perimeter
				plt.perimeter++
			}
			if seen[p1] {
				continue
			}
			if v == plant {
				q.Push(p1)
				seen[p1] = true
			}
		}
	}
	return plt.area * plt.perimeter
}

func check(grid [][]rune, p point) bool {
	return p.x >= 0 && p.x < len(grid[0]) && p.y >= 0 && p.y < len(grid)
}
