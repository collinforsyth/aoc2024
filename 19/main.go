package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"iter"
	"strings"
	"unicode"
)

func main() {
	input, err := util.ReadInput("19/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	dict, words := parseInput(input.Lines())
	fmt.Println("Part 1: ", partOne(dict, words))
	fmt.Println("Part 2: ", partTwo(dict, words))
}

func partOne(dict []string, words []string) int {
	t := &Trie{children: make(map[rune]*Trie)}
	for _, word := range dict {
		t.Insert(word)
	}
	count := 0
	for _, word := range words {
		if dfs(word, t) {
			count++
		}
	}
	return count
}

func partTwo(dict []string, words []string) int {
	return 0
}

func dfs(s string, t *Trie) bool {
	if len(s) == 0 {
		return true
	}
	for i := 1; i <= len(s); i++ {
		if t.Search(s[:i]) {
			if dfs(s[i:], t) {
				return true
			}
		}
	}
	return false
}

func parseInput(input iter.Seq2[int, string]) ([]string, []string) {
	var (
		dict  []string
		words []string
	)
	for i, line := range input {
		switch i {
		case 0:
			dict = strings.FieldsFunc(line, func(r rune) bool {
				return r == ',' || unicode.IsSpace(r)
			})
		case 1:
			continue
		default:
			words = append(words, line)
		}
	}
	return dict, words
}

type Trie struct {
	children map[rune]*Trie
	end      bool
}

func (t *Trie) Insert(s string) {
	n := t
	for _, r := range s {
		if _, ok := n.children[r]; !ok {
			n.children[r] = &Trie{children: make(map[rune]*Trie)}
		}
		n = n.children[r]
	}
	n.end = true
}
func (t *Trie) Search(s string) bool {
	n := t
	for _, r := range s {
		if _, ok := n.children[r]; !ok {
			return false
		}
		n = n.children[r]
	}
	return n.end
}
func (t *Trie) All(s string) []string {
	var res []string
	n := t
	for _, r := range s {
		if _, ok := n.children[r]; !ok {
			return res
		}
		if n.end {
			res = append(res, s)
		}
	}
	return res
}
