package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (input []int) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		input = append(input, aoc.StrToInt(line))
	}

	return input
}

func process(input []int) int {
	count := 0
	permutations(input, 150, &count)

	return count
}

func permutations(containers []int, volumeLeft int, counter *int) {
	for i, container := range containers {
		if container > volumeLeft {
			continue
		}
		if container == volumeLeft {
			*counter++
		}
		permutations(containers[i+1:], volumeLeft-container, counter)
	}
}
