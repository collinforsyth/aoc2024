package main

import (
	"bufio"
	"collinforsyth/aoc2024/util"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	inputFile, err := os.Open("02/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()
	input := parseInput(inputFile)
	partOne(input)
	partTwo(input)
}

func parseInput(inputFile io.Reader) [][]int {
	res := make([][]int, 0)
	sc := bufio.NewScanner(inputFile)
	for sc.Scan() {
		vals := strings.Fields(sc.Text())
		tmp := make([]int, len(vals))
		for i := range vals {
			tmp[i] = util.MustAtoi(vals[i])
		}
		res = append(res, tmp)
	}
	return res
}

func partOne(reports [][]int) {
	sum := 0
	for _, report := range reports {
		if isSafe(report) {
			sum++
		}
	}
	fmt.Println("Part 1: ", sum)
}

func partTwo(reports [][]int) {
	sum := 0
	for _, report := range reports {
		if isSafe(report) {
			sum++
		} else {
			// brute force all possible combinations
			for i := range report {
				tmp := slices.Delete(slices.Clone(report), i, i+1)
				if isSafe(tmp) {
					sum++
					break
				}
			}
		}
	}
	fmt.Println("Part 2: ", sum)
}

func isSafe(levels []int) bool {
	cmp := strictlyIncreasing
	// peek to see if the levels should be in ascending or descending order
	if levels[0]-levels[1] > 0 {
		cmp = strictlyDecreasing
	}
	for i := len(levels) - 1; i > 0; i-- {
		if !safeDiff(levels[i-1], levels[i]) || !cmp(levels[i-1], levels[i]) {
			return false
		}
	}
	return true
}

func safeDiff(i, j int) bool {
	diff := util.Abs(i - j)
	return diff >= 1 && diff <= 3
}

func strictlyIncreasing(i, j int) bool {
	return i < j
}

func strictlyDecreasing(i, j int) bool {
	return i > j
}
