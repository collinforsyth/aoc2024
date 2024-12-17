package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	input, err := util.ReadInput("17/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	opcodes, registers := parseInput(input.String())
	fmt.Println("Part 1: ", output(partOne(opcodes, registers)))
}

func parseInput(input string) ([]int, []int) {
	register := regexp.MustCompile(`Register (.): (\d+)`)
	program := regexp.MustCompile(`Program: (([\d+][,]?)+)`)

	rMatches := register.FindAllStringSubmatch(input, -1)
	pMatches := program.FindAllStringSubmatch(input, -1)
	registers := make([]int, 0)
	for _, r := range rMatches {
		registers = append(registers, util.MustAtoi(r[2]))
	}
	opCodes := make([]int, 0)
	fields := strings.FieldsFunc(pMatches[0][1], func(r rune) bool {
		return r == ','
	})
	for _, oc := range fields {
		opCodes = append(opCodes, util.MustAtoi(oc))

	}
	return registers, opCodes
}

// So, the program 0,1,2,3 would run the
// instruction whose opcode is 0 and pass it
// the operand 1, then run the instruction having
// opcode 2 and pass it the operand 3, then halt.
func partOne(registers []int, opCodes []int) []int {
	stdOut := make([]int, 0)
	pc := 0
	for {
		// if we're ever at the end, halt
		// NOTE: assuming end is i+1 since we need 2
		if pc >= len(opCodes)-1 {
			break
		}
		switch opCodes[pc] {
		case 0:
			d := int(math.Pow(2, float64(combo(opCodes[pc+1], registers))))
			n := registers[0]
			registers[0] = n / d
		case 1:
			registers[1] = registers[1] ^ opCodes[pc+1]
		case 2:
			c := combo(opCodes[pc+1], registers)
			registers[1] = c % 8
		case 3:
			if registers[0] != 0 {
				pc = opCodes[pc+1]
				continue
			}
		case 4:
			registers[1] = registers[1] ^ registers[2]
		case 5:
			c := combo(opCodes[pc+1], registers) % 8
			stdOut = append(stdOut, c)
		case 6:
			d := int(math.Pow(2, float64(combo(opCodes[pc+1], registers))))
			n := registers[0]
			registers[1] = n / d
		case 7:
			d := int(math.Pow(2, float64(combo(opCodes[pc+1], registers))))
			n := registers[0]
			registers[2] = n / d
		}
		pc += 2
	}
	return stdOut
}

func combo(i int, registers []int) int {
	switch i {
	case 0, 1, 2, 3:
		return i
	case 4:
		return registers[0]
	case 5:
		return registers[1]
	case 6:
		return registers[2]
	case 7:
		panic("reserved")
	default:
		panic("unknownw combo")
	}
}

func output(a []int) string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = fmt.Sprint(v)
	}
	return strings.Join(s, ",")
}
