package main

import (
	"fmt"
	"io/ioutil"
)

func partOne(input string) int {
	return 1
}

func partTwo(input string) int {
	return 1
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
