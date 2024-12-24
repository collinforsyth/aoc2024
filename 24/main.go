package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := util.ReadInput("24/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	wires, gates := parseInput(input)
	fmt.Println("Part One: ", partOne(wires, gates))
	fmt.Println("Part Two: ", partTwo(wires, gates, 4))
}

type Operator int

const (
	AND Operator = iota
	OR
	XOR
)

type gate struct {
	a, b   string
	op     Operator
	output string
}

func parseInput(input *util.Input) (map[string]uint, []gate) {
	gates := make([]gate, 0)
	wires := make(map[string]uint)
	for _, line := range input.Lines() {
		f := strings.Fields(line)
		switch len(f) {
		case 2:
			s, _ := strings.CutSuffix(f[0], ":")
			b, _ := strconv.ParseUint(f[1], 10, 64)
			wires[s] = uint(b)
		case 5:
			g := gate{a: f[0], b: f[2], output: f[4]}
			switch f[1] {
			case "AND":
				g.op = AND
			case "OR":
				g.op = OR
			case "XOR":
				g.op = XOR
			default:
				panic("unknown operator")
			}
			gates = append(gates, g)
		default:
			continue
		}
	}
	return wires, gates
}

func partOne(wires map[string]uint, gates []gate) uint {
	m := make(map[string][]string)
	for _, g := range gates {
		// for each input, they have a direction to output
		m[g.a] = append(m[g.a], g.output)
		m[g.b] = append(m[g.b], g.output)
	}

	stack := make([]string, 0)
	visited := make(map[string]bool)
	for node := range m {
		if !visited[node] {
			topologicalSortUtil(m, node, visited, &stack)
		}
	}
	// assign precedence to each gate
	precedences := make(map[string]int)
	for i, n := range stack {
		precedences[n] = i
	}
	slices.SortFunc(gates, func(a, b gate) int {
		return precedences[b.output] - precedences[a.output]
	})
	for _, g := range gates {
		a, b := wires[g.a], wires[g.b]
		switch g.op {
		case AND:
			wires[g.output] = a & b
		case OR:
			wires[g.output] = a | b
		case XOR:
			wires[g.output] = a ^ b
		}
	}
	output := register(wires, "z")
	return output
}

func topologicalSortUtil(g map[string][]string, n string, visited map[string]bool, stack *[]string) {
	visited[n] = true
	for _, neighbor := range g[n] {
		if !visited[neighbor] {
			topologicalSortUtil(g, neighbor, visited, stack)
		}
	}
	*stack = append(*stack, n)
}

func partTwo(wires map[string]uint, gates []gate, swaps int) string {
	return ""
}

func register(wires map[string]uint, prefix string) uint {
	return toUint(wires, vals(wires, prefix))
}

func vals(wires map[string]uint, prefix string) []string {
	vals := []string{}
	for w := range wires {
		if strings.HasPrefix(w, prefix) {
			vals = append(vals, w)
		}
	}
	slices.Sort(vals)
	return vals
}

func toUint(wires map[string]uint, vals []string) uint {
	var output uint = 0
	for i, w := range vals {
		v := wires[w]
		output = output | (v << uint(i))
	}
	return output
}
