package main

import (
	"collinforsyth/aoc2024/util"
	"fmt"
	"slices"
)

func main() {
	input, err := util.ReadInput("09/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1: ", partOne(input.String()))
}

type layout struct {
	id    int
	files []int
	free  []int
}

func partOne(input string) int {
	diskMap := make([]layout, 0)
	i := 0
	for i < len(input)-1 {
		id := i / 2
		files := slices.Repeat([]int{id}, runeToDigit(input[i]))
		free := slices.Repeat([]int{-1}, runeToDigit(input[i+1]))
		l := layout{
			id:    id,
			files: files,
			free:  free,
		}
		diskMap = append(diskMap, l)
		i += 2
	}
	if len(input)%2 != 0 {
		files := slices.Repeat([]int{i / 2}, runeToDigit(input[i]))
		l := layout{
			id:    i / 2,
			files: files,
		}
		diskMap = append(diskMap, l)
	}

	i = 0
	j := len(diskMap) - 1
	k, l := 0, len(diskMap[j].files)-1
	for i < j {
		if l < 0 {
			j--
			l = len(diskMap[j].files) - 1
			continue
		}
		if k >= len(diskMap[i].free) {
			i++
			k = 0
			continue
		}
		// find the next free
		for k < len(diskMap[i].free) && diskMap[i].free[k] != -1 {
			k++
			continue
		}
		// find the next swap
		for l > 0 && diskMap[j].files[l] == -1 {
			l--
			continue
		}
		// swap
		diskMap[i].free[k], diskMap[j].files[l] = diskMap[j].files[l], diskMap[i].free[k]
		k++
		l--
	}

	i = 0
	sum := 0
	for _, d := range diskMap {
		for _, f := range d.files {
			if f == -1 {
				continue
			}
			sum += (f * i)
			i++
		}
		for _, f := range d.free {
			if f == -1 {
				continue
			}
			sum += (f * i)
			i++
		}
	}
	return sum
}

func runeToDigit(r byte) int {
	return int(r - '0')
}
