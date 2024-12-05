package util

import (
	"bufio"
	"bytes"
	"iter"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Input struct {
	input *bytes.Buffer
}

func ReadInput(file string) (*Input, error) {
	b, err := os.ReadFile(file)
	if err != nil {
	}
	buf := bytes.NewBuffer(b)
	return &Input{input: buf}, nil
}

func FromBytes(b []byte) *Input {
	return &Input{input: bytes.NewBuffer(b)}
}

func (i *Input) String() string {
	return i.input.String()
}

func (i *Input) Bytes() []byte {
	return i.input.Bytes()
}

func (i *Input) Lines() iter.Seq[string] {
	sc := bufio.NewScanner(strings.NewReader(i.input.String()))
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
