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

func (i *Input) Clone() *Input {
	return &Input{input: bytes.NewBuffer(i.input.Bytes())}
}

func (i *Input) Runes() [][]rune {
	b := make([][]rune, 0)
	sc := bufio.NewScanner(bytes.NewReader(i.input.Bytes()))
	for sc.Scan() {
		b = append(b, []rune(string(sc.Bytes())))
	}
	return b
}

func (i *Input) Lines() iter.Seq2[int, string] {
	sc := bufio.NewScanner(strings.NewReader(i.input.String()))
	j := 0
	return func(yield func(int, string) bool) {
		for sc.Scan() {
			if !yield(j, sc.Text()) {
				break
			}
			j++
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

func MustFloat(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return n
}

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() (T, bool) {
	var v T
	if len(*q) == 0 {
		return v, false
	}
	v = (*q)[0]
	*q = (*q)[1:]
	return v, true
}

func (q Queue[T]) Len() int {
	return len(q)
}
