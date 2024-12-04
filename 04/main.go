package main

import (
	"bufio"
	"bytes"
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
)

func main() {
	input, err := util.ReadInput("04/input.txt")
	if err != nil {
		log.Println(err)
		return
	}
	parsed := parseInput(input.Bytes())
	fmt.Println("Part 1: ", partOne(parsed))
}

func parseInput(input []byte) [][]byte {
	b := make([][]byte, 0)
	sc := bufio.NewScanner(bytes.NewReader(input))
	for sc.Scan() {
		b = append(b, sc.Bytes())
	}
	return b
}

type point struct {
	x, y int
}

var directions = []point{
	{-1, 0},  // left
	{1, 0},   // right
	{0, -1},  // up
	{0, 1},   // down
	{-1, -1}, // up-left
	{1, 1},   // up-right
	{-1, 1},  // down-left
	{1, -1},  // down-right
}

func partOne(input [][]byte) int {
	sum := 0
	for i := range input {
		for j := range input[i] {
			if input[i][j] != 'X' {
				continue
			}
			// for each direction - up, down, left, right, up-left, up-right, down-left, down-right
			for _, d := range directions {
				if dfs(input, i, j, []byte{'X', 'M', 'A', 'S'}, 0, d) {
					sum++
				}

			}
		}
	}
	return sum
}

func dfs(input [][]byte, i, j int, word []byte, k int, d point) bool {
	// clamp bounds
	if i < 0 || j < 0 || i >= len(input) || j >= len(input[i]) {
		return false
	}
	if input[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	// temporarily mark as visited
	return dfs(input, i+d.y, j+d.x, word, k+1, d)
}
