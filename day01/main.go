package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func caloriesPerElf(input string) []int {
	var elves []int

	sum := 0

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			elves = append(elves, sum)
			sum = 0
			continue
		}
		i, _ := strconv.Atoi(line)
		sum += i
	}

	if sum > 0 {
		elves = append(elves, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	return elves
}

func partOne(elves []int) int {
	return elves[0]
}

func partTwo(elves []int) int {
	return elves[0] + elves[1] + elves[2]
}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	elves := caloriesPerElf(string(data))
	fmt.Printf("Part one: %d\n", partOne(elves))
	fmt.Printf("Part two: %d\n", partTwo(elves))
}
