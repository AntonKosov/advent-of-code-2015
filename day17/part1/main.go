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
	counter := map[int]int{} // count containers -> count variants
	permutations(input, 150, 0, counter)
	minContainers := len(input)
	variants := 0
	for countainers, count := range counter {
		if countainers < minContainers {
			minContainers = countainers
			variants = count
		}
	}

	return variants
}

func permutations(containers []int, volumeLeft int, usedContainers int, counter map[int]int) {
	for i, container := range containers {
		if container > volumeLeft {
			continue
		}
		if container == volumeLeft {
			counter[usedContainers+1]++
		}
		permutations(containers[i+1:], volumeLeft-container, usedContainers+1, counter)
	}
}
