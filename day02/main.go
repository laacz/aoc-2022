package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func partOne(input string) int {
	score := 0

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		elf_move := -('A' - int(line[0]) - 1)
		my_move := -('X' - int(line[2]) - 1)

		score += my_move

		if my_move-elf_move == 1 || (my_move == 1 && elf_move == 3) {
			score += 6
		} else if my_move == elf_move {
			score += 3
		}
	}

	return score
}

func partTwo(input string) int {
	score := 0
	my_move := 0

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		elf_move := -('A' - int(line[0]) - 1)

		switch line[2:3] {
		case "X": // must lose
			if elf_move == 1 {
				my_move = 3
			} else {
				my_move = elf_move - 1
			}
		case "Z": // must win
			if elf_move == 3 {
				my_move = 1
			} else {
				my_move = elf_move + 1
			}
			score += 6
			break
		case "Y": // myst draw
			my_move = elf_move
			score += 3
		}

		score += my_move
	}

	return score
}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
