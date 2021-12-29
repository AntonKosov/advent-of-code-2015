package main

import (
	"fmt"
	"regexp"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type properties []int

func read() (input []properties) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	re := regexp.MustCompile(`-?\d+`)

	for _, line := range lines {
		parts := re.FindAll([]byte(line), -1)
		var props properties
		for i := 0; i < 4; i++ {
			props = append(props, aoc.StrToInt(string(parts[i])))
		}
		input = append(input, props)
	}

	return input
}

func process(input []properties) int {
	volumes := make([]int, len(input))

	max := 0
	permutations(0, volumes, 100, func() {
		total := 1
		for pi := 0; pi < len(input[0]); pi++ {
			st := 0
			for si, props := range input {
				st += props[pi] * volumes[si]
			}
			if st <= 0 {
				return
			}
			total *= st
		}
		max = aoc.Max(max, total)
	})

	return max
}

func permutations(index int, volumes []int, volumeLeft int, ready func()) {
	if index == len(volumes)-1 {
		volumes[index] = volumeLeft
		ready()
		return
	}

	for i := 0; i <= volumeLeft; i++ {
		volumes[index] = i
		permutations(index+1, volumes, volumeLeft-i, ready)
	}
}
