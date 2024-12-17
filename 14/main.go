package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"iter"
	"regexp"
)

func main() {
	input, err := util.ReadInput("14/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	robots := parseInput(input.Lines())
	fmt.Println("Part 1: ", partOne(101, 103, robots))
}

type robot struct {
	p point
	v point
}

type point struct {
	x, y int
}

func parseInput(input iter.Seq2[int, string]) []robot {
	robots := make([]robot, 0)
	re := regexp.MustCompile(`(-?\d+),(-?\d+)`)
	for _, s := range input {
		matches := re.FindAllStringSubmatch(s, -1)
		r := robot{
			p: point{x: util.MustAtoi(matches[0][1]), y: util.MustAtoi(matches[0][2])},
			v: point{x: util.MustAtoi(matches[1][1]), y: util.MustAtoi(matches[1][2])},
		}
		robots = append(robots, r)
	}
	return robots
}

func partOne(x, y int, robots []robot) int {
	// map[robotID] => current position
	memo := make(map[int]point)
	for i := range robots {
		memo[i] = robots[i].p
	}
	for range 100 {
		for i, r := range robots {
			memo[i] = next(x, y, memo[i], r.v)
		}
	}
	return safetyFactor(x, y, memo)
}

func next(x, y int, p, v point) point {
	newP := point{x: p.x + v.x, y: p.y + v.y}
	if newP.x < 0 {
		newP.x = x + newP.x
	} else if newP.x >= x {
		newP.x = newP.x - x
	}
	if newP.y < 0 {
		newP.y = y + newP.y
	} else if newP.y >= y {
		newP.y = newP.y - y
	}
	return newP
}

func safetyFactor(x, y int, r map[int]point) int {
	midX := x / 2
	midY := y / 2
	var quadrants = [4]int{0, 0, 0, 0}
	for _, p := range r {
		if p.x == midX || p.y == midY {
			continue
		} else if p.x < midX && p.y < midY {
			quadrants[0]++
		} else if p.x > midX && p.y < midY {
			quadrants[1]++
		} else if p.x < midX && p.y > midY {
			quadrants[2]++
		} else {
			quadrants[3]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}
