package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"iter"
	"math"
	"os"
	"regexp"
)

func main() {
	input, err := util.ReadInput("13/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	machines := parseInput(input.Lines(), 0)
	fmt.Println("Part 1:", partOne(machines))
	machines = parseInput(input.Lines(), 10000000000000)
	fmt.Println("Part 2:", partOne(machines))
}

type machine struct {
	a     point
	b     point
	prize point
}

func (m machine) String() string {
	return fmt.Sprintf("a:%v, b:%v, p:%v", m.a, m.b, m.prize)
}

type point struct {
	x, y float64
}

func (p point) String() string {
	return fmt.Sprintf("{%f, %f}", p.x, p.y)
}

func parseInput(input iter.Seq2[int, string], prizeOffset float64) []machine {
	machines := make([]machine, 0)
	tmp := machine{}
	j := 0
	for _, s := range input {
		if s == "" {
			machines, j = append(machines, tmp), 0
			continue
		}
		switch j % 3 {
		case 0:
			tmp.a = parseXY(s)
		case 1:
			tmp.b = parseXY(s)
		case 2:
			p := parseXY(s)
			tmp.prize = point{p.x + prizeOffset, p.y + prizeOffset}
		}
		j++
	}
	machines = append(machines, tmp)
	return machines
}

func parseXY(s string) point {
	r := regexp.MustCompile(`X(\+|=)(\d+), Y(\+|=)(\d+)`)
	p := r.FindStringSubmatch(s)
	return point{
		x: util.MustFloat(p[2]),
		y: util.MustFloat(p[4]),
	}
}

func partOne(machines []machine) int {
	sum := 0
	for _, m := range machines {
		m := gaussianElimination([][]float64{
			{m.a.x, m.b.x},
			{m.a.y, m.b.y},
		},
			[]float64{m.prize.x, m.prize.y},
		)
		if almostEqual(m[0]) && almostEqual(m[1]) {
			sum += int(math.Round(m[0]))*3 + int(math.Round(m[1]))
		}
	}
	return sum
}

// https://en.wikipedia.org/wiki/Gaussian_elimination#Pseudocode
func gaussianElimination(A [][]float64, b []float64) []float64 {
	n := len(A)
	if n != len(b) {
		panic("matrix A and vector b must have the same number of rows")
	}

	// Forward elimination
	for i := 0; i < n-1; i++ {
		// Partial pivoting
		maxRow := i
		for j := i + 1; j < n; j++ {
			if math.Abs(A[j][i]) > math.Abs(A[maxRow][i]) {
				maxRow = j
			}
		}
		if maxRow != i {
			// swap rows i and maxRow
			A[i], A[maxRow] = A[maxRow], A[i]
			b[i], b[maxRow] = b[maxRow], b[i]
		}

		// Elimination
		for j := i + 1; j < n; j++ {
			factor := A[j][i] / A[i][i]
			for k := i; k < n; k++ {
				A[j][k] -= factor * A[i][k]
			}
			b[j] -= factor * b[i]
		}
	}

	// Back substitution
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += A[i][j] * x[j]
		}
		x[i] = (b[i] - sum) / A[i][i]
	}

	return x
}

const epsilon = 1e-3

func almostEqual(a float64) bool {
	return math.Abs(a-math.Round(a)) < epsilon
}
