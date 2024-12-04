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
	{-1, -1}, // left-up
	{1, 1},   // right-down
	{-1, 1},  // left-down
	{1, -1},  // right-down
}

func partOne(input [][]byte) int {
	sum := 0
	word := []byte{'X', 'M', 'A', 'S'}
	for i := range input {
		for j := range input[i] {
			if input[i][j] != 'X' {
				continue
			}
			for _, d := range directions {
				if search(input, word, i, j, d) {
					sum++
				}
			}
		}
	}
	return sum
}

func search(input [][]byte, word []byte, i, j int, d point) bool {
	k := 0
	for input[i][j] == word[k] {
		if k == len(word)-1 {
			return true
		}
		i, j = i+d.y, j+d.x
		if !check(input, i, j) {
			return false
		}
		k++
	}
	return false
}

func check(input [][]byte, i, j int) bool {
	if i < 0 || j < 0 || i >= len(input) || j >= len(input[i]) {
		return false
	}
	return true
}
