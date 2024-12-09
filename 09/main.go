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
	diskMap := parseInput(input.String())
	fmt.Println("Part 1: ", partOne(diskMap))
	diskMap = parseInput(input.String())
	fmt.Println("Part 2: ", partTwo(diskMap))
}

type layout struct {
	id    int
	files []int
	free  []int
}

func parseInput(input string) []layout {
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

	return diskMap
}

func partOne(diskMap []layout) int {
	i := 0
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

	return checksum(diskMap)
}

func partTwo(diskMap []layout) int {
	for j := len(diskMap) - 1; j >= 0; j-- {
		// if we can't fit the files into the free space,
		// move to the next disk
		for i := 0; i < j; i++ {
			free := free(diskMap[i].free)
			if len(diskMap[j].files) > free {
				continue
			}
			l := len(diskMap[i].free) - free
			for k := 0; k < len(diskMap[j].files); k++ {
				diskMap[i].free[l], diskMap[j].files[k] = diskMap[j].files[k], diskMap[i].free[l]
				l++
			}
			break
		}
	}
	return checksum(diskMap)
}

func free(l []int) int {
	c := 0
	for _, f := range l {
		if f == -1 {
			c++
		}
	}
	return c
}

func checksum(diskMap []layout) int {
	i := 0
	sum := 0
	for _, l := range diskMap {
		for _, f := range l.files {
			if f == -1 {
				i++
				continue
			}
			sum += (f * i)
			i++
		}
		for _, f := range l.free {
			if f == -1 {
				i++
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
