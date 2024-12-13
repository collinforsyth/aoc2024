package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"iter"
	"log"
	"math"
	"strings"
	"unicode"
)

func main() {
	input, err := util.ReadInput("07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parsed := parseInput(input)
	fmt.Println("Part 1: ", partOne(parsed))
	fmt.Println("Part 2: ", partTwo(parsed))
}

type equation struct {
	value   int
	factors []int
}

func parseInput(input *util.Input) []equation {
	equations := make([]equation, 0)
	for _, line := range input.Lines() {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return unicode.IsSpace(r) || r == ':'
		})
		factors := make([]int, 0)
		for _, f := range fields[1:] {
			factors = append(factors, util.MustAtoi(f))
		}
		equations = append(equations, equation{
			value:   util.MustAtoi(fields[0]),
			factors: factors,
		})
	}
	return equations
}

func partOne(input []equation) int {
	res := 0
	for _, eq := range input {
		for c := range product([]string{"+", "*"}, len(eq.factors)-1) {
			sum := eq.factors[0]
			for i := 1; i < len(eq.factors); i++ {
				switch c[i-1] {
				case "+":
					sum += eq.factors[i]
				case "*":
					sum *= eq.factors[i]
				}
			}
			if sum == eq.value {
				res += eq.value
				break
			}
		}
	}
	return res
}

func partTwo(input []equation) int {
	res := 0
	for _, eq := range input {
		for c := range product([]string{"+", "*", "||"}, len(eq.factors)-1) {
			sum := eq.factors[0]
			for i := 1; i < len(eq.factors); i++ {
				switch c[i-1] {
				case "+":
					sum += eq.factors[i]
				case "*":
					sum *= eq.factors[i]
				case "||":
					// sum = util.MustAtoi(strconv.Itoa(sum) + strconv.Itoa(eq.factors[i]))
					sum = concat(sum, eq.factors[i])
				}
			}
			if sum == eq.value {
				res += eq.value
				break
			}
		}
	}
	return res
}

func concat(a, b int) int {
	i, tmp := 0, b
	for tmp != 0 {
		tmp /= 10
		i++
	}
	return a*int(math.Pow10(i)) + b
}

func product[T any](vals []T, k int) iter.Seq[[]T] {
	indexes := make([]int, k)
	p := make([]T, k)
	// a = a[:0] clear slice
	return func(yield func([]T) bool) {
		for indexes != nil {
			for i, x := range indexes {
				p[i] = vals[x]
			}
			for i := len(indexes) - 1; i >= 0; i-- {
				indexes[i]++
				if indexes[i] < len(vals) {
					break
				}
				indexes[i] = 0
				if i <= 0 {
					indexes = nil
					break
				}
			}
			if !yield(p) {
				return
			}
		}
	}
}
