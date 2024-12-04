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

func parseInput(input []byte) [][]rune {
	b := make([][]rune, 0)
	sc := bufio.NewScanner(bytes.NewReader(input))
	for sc.Scan() {
		b = append(b, []rune(string(sc.Bytes())))
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

func partOne(input [][]rune) int {
	sum := 0
	word := []rune{'X', 'M', 'A', 'S'}
	for i := range input {
		for j := range input[i] {
			for _, d := range directions {
				if search(input, word, i, j, d) {
					sum++
				}
			}
		}
	}
	return sum
}

func search(input [][]rune, word []rune, i, j int, d point) bool {
	for k := 0; k < len(word); k++ {
		if !check(input, i, j) || input[i][j] != word[k] {
			return false
		}
		i += d.y
		j += d.x
	}
	return true
}

func check(input [][]rune, i, j int) bool {
	if i < 0 || j < 0 || i >= len(input) || j >= len(input[i]) {
		return false
	}
	return true
}