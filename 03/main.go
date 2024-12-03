package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	bytes, err := os.ReadFile("03/input.txt")
	input := string(bytes)
	if err != nil {
		log.Println(err)
	}
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	sum := 0
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	res := re.FindAllStringSubmatch(input, -1)
	for _, r := range res {
		sum += util.MustAtoi(r[1]) * util.MustAtoi(r[2])
	}
	fmt.Println("Part 1: ", sum)
}

func partTwo(input string) {
	sum := 0
	re := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\)|(do(n't)?))`)
	res := re.FindAllStringSubmatch(input, -1)
	direction := true
	for _, r := range res {
		switch r[0] {
		case "don't":
			direction = false
		case "do":
			direction = true
		default:
			if direction {
				sum += util.MustAtoi(r[2]) * util.MustAtoi(r[3])
			}
		}
	}
	fmt.Println("Part 2: ", sum)
}
