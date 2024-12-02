package main

import (
	"bufio"
	"collinforsyth/aoc2024/util"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("02/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := parseInput(inputFile)
	if err := partOne(input); err != nil {
		fmt.Println(err)
		return
	}
	if err := partTwo(input); err != nil {
		fmt.Println(err)
		return
	}

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

func partOne(reports [][]int) error {
	sum := 0
	for _, report := range reports {
		if isSafe(report, 0) {
			sum++
		}
	}
	fmt.Println("Part 1: ", sum)
	return nil
}

func partTwo(reports [][]int) error {
	sum := 0
	for _, report := range reports {
		if isSafe(report, 1) {
			sum++
		}
	}
	fmt.Println("Part 1: ", sum)
	return nil
}

func isSafe(levels []int, toleration int) bool {
	cmp := strictlyIncreasing
	// peek to see if the levels should be in ascending or descending order
	if levels[0]-levels[1] < 0 {
		cmp = strictlyDecreasing
	}
	badLevelCount := 0
	for i := len(levels) - 1; i > 0; i-- {
		if !safeDiff(levels[i-1], levels[i]) {
			badLevelCount++
			continue
		}
		if !cmp(levels[i-1], levels[i]) {
			badLevelCount++
			continue
		}
	}
	// return early if possible
	return badLevelCount <= toleration
}

func safeDiff(i, j int) bool {
	diff := util.Abs(i - j)
	return diff >= 1 && diff <= 3
}

func strictlyIncreasing(i, j int) bool {
	return i > j
}

func strictlyDecreasing(i, j int) bool {
	return i < j
}
