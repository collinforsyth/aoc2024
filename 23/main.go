package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"iter"
	"slices"
	"strings"
)

func main() {
	input, err := util.ReadInput("23/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	p := parseInput(input)
	fmt.Println("Part One: ", partOne(p))
	fmt.Println("Part Two: ", partTwo(p))
}

func parseInput(i *util.Input) map[string][]string {
	m := make(map[string][]string)
	for _, line := range i.Lines() {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-'
		})
		c1, c2 := fields[0], fields[1]
		m[c1] = append(m[c1], c2)
		m[c2] = append(m[c2], c1)
	}
	return m
}

func partOne(p map[string][]string) int {
	cycles := make([][]string, 0)
	visited := make(map[string]bool)
	for node := range p {
		dfs1(p, visited, node, node, []string{}, &cycles)
		visited[node] = true
	}
	for _, cycle := range cycles {
		slices.Sort(cycle)
	}
	slices.SortFunc(cycles, func(a, b []string) int {
		return slices.Compare(a, b)
	})
	cycles = slices.CompactFunc(cycles, func(a, b []string) bool {
		return slices.Compare(a, b) == 0
	})
	count := 0
	for _, cycle := range cycles {
		for _, c := range cycle {
			if strings.HasPrefix(c, "t") {
				count++
				break
			}
		}
	}
	return count
}

func dfs1(p map[string][]string, visited map[string]bool, start, node string, path []string, cycles *[][]string) {
	if len(path) == 3 && node == start {
		*cycles = append(*cycles, path)
	}
	if len(path) >= 3 || visited[node] {
		return
	}
	for _, neighbor := range p[node] {
		// if already visited
		if slices.Contains(path, neighbor) {
			continue
		}
		dfs1(p, visited, start, neighbor, append(path, neighbor), cycles)
	}
}

func partTwo(p map[string][]string) string {
	var password []string
	for n, neighbors := range p {
		for i := 0; i < len(neighbors); i++ {
			if len(neighbors[i:]) < len(password) {
				continue
			}
			var allConnected = true
			for pair := range pairs(neighbors[i:]) {
				if !slices.Contains(p[pair[0]], pair[1]) || !slices.Contains(p[pair[1]], pair[0]) {
					allConnected = false
					break
				}
			}
			if allConnected {
				password = append(neighbors[i:], n)
			}
		}
	}
	slices.Sort(password)
	result := strings.Join(password, ",")
	return result
}

func pairs[T any](slice []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for i := 0; i < len(slice)-1; i++ {
			for j := i + 1; j < len(slice); j++ {
				p := []T{slice[i], slice[j]}
				if !yield(p) {
					return
				}
			}
		}
	}
}
