package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func main() {
	input, err := util.ReadInput("17/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	registers, program := parseInput(input.String())
	fmt.Println("Part 1: ", output(partOne(registers, program)))
	fmt.Println("Part 2: ", partTwo(program))
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

func partOne(registers []int, opCodes []int) []int {
	c := computer{registers: registers, program: opCodes}
	c.run()
	return c.stdout
}

func partTwo(program []int) int {
	i := 0
	for p := len(program) - 1; p >= 0; p-- {
		i <<= 3
		for !slices.Equal(run([]int{i, 0, 0}, program), program[p:]) {
			i++
		}
		fmt.Println(i, program[p:])
	}
	return i
}

type computer struct {
	registers []int
	program   []int
	pc        int
	stdout    []int
}

func run(registers []int, program []int) []int {
	c := computer{registers: registers, program: program}
	c.run()
	return c.stdout
}

func (c *computer) run() {
	for c.pc < len(c.program) {
		c.exec()
	}
}

func (c *computer) exec() {
	operand := c.program[c.pc+1]
	switch c.program[c.pc] {
	case 0:
		c.registers[0] = c.registers[0] >> combo(operand, c.registers)
	case 1:
		c.registers[1] = c.registers[1] ^ operand
	case 2:
		c.registers[1] = combo(operand, c.registers) % 8
	case 3:
		if c.registers[0] != 0 {
			c.pc = operand
			return
		}
	case 4:
		c.registers[1] = c.registers[1] ^ c.registers[2]
	case 5:
		c.stdout = append(c.stdout, combo(operand, c.registers)%8)
	case 6:
		c.registers[1] = c.registers[0] >> combo(operand, c.registers)
	case 7:
		c.registers[2] = c.registers[0] >> combo(operand, c.registers)
	}
	c.pc += 2
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
		panic("unknown combo")
	}
}

func output(a []int) string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = fmt.Sprint(v)
	}
	return strings.Join(s, ",")
}
