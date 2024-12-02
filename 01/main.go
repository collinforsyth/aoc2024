package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"collinforsyth/aoc2024/util"
)

func main() {
	input, err := os.Open("01/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()
	l1, l2 := parseInput(input)
	if err := partOne(l1, l2); err != nil {
		fmt.Println(err)
		return
	}
	if err := partTwo(l1, l2); err != nil {
		fmt.Println(err)
		return
	}

}

func parseInput(inputFile io.Reader) ([]int, []int) {
	sc := bufio.NewScanner(inputFile)
	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for sc.Scan() {
		vals := strings.Fields(sc.Text())
		n1, n2 := util.MustAtoi(vals[0]), util.MustAtoi(vals[1])
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	return l1, l2
}

func partOne(l1, l2 []int) error {
	sum := 0
	slices.Sort(l1)
	slices.Sort(l2)
	for i := range l1 {
		sum += util.Abs(l1[i] - l2[i])
	}
	fmt.Println("Part 1: ", sum)
	return nil
}

func partTwo(l1, l2 []int) error {
	counts := make(map[int]int)
	sum := 0
	for _, n := range l2 {
		counts[n]++
	}
	for _, n := range l1 {
		if c, ok := counts[n]; ok {
			sum += n * c
		}
	}
	fmt.Println("Part 2: ", sum)
	return nil
}
