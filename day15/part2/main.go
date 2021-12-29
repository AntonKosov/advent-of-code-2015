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

type ingredient struct {
	properties []int
	calories   int
}

func read() (input []ingredient) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	re := regexp.MustCompile(`-?\d+`)

	for _, line := range lines {
		parts := re.FindAll([]byte(line), -1)
		var ing ingredient
		for i := 0; i < 4; i++ {
			ing.properties = append(ing.properties, aoc.StrToInt(string(parts[i])))
		}
		ing.calories = aoc.StrToInt(string(parts[4]))
		input = append(input, ing)
	}

	return input
}

func process(input []ingredient) int {
	volumes := make([]int, len(input))

	max := 0
	permutations(0, volumes, 100, func() {
		calories := 0
		for i, ing := range input {
			calories += ing.calories * volumes[i]
		}
		if calories != 500 {
			return
		}
		total := 1
		for pi := 0; pi < len(input[0].properties); pi++ {
			st := 0
			for si, ing := range input {
				st += ing.properties[pi] * volumes[si]
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
