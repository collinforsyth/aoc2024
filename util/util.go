package util

import (
	"bufio"
	"iter"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Input struct {
	input string
}

func ReadInput(file string) (*Input, error) {
	f, err := os.ReadFile(file)
	if err != nil {
	}
	return &Input{input: string(f)}, nil
}

func (i *Input) String() string {
	return i.input
}

func (i *Input) Lines() iter.Seq[string] {
	sc := bufio.NewScanner(strings.NewReader(i.input))
	return func(yield func(string) bool) {
		for sc.Scan() {
			if !yield(sc.Text()) {
				break
			}
		}
	}
}

func Abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
