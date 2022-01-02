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
	for i := len(lines) - 2; i >= 0; i-- {
		input = append(input, aoc.StrToInt(lines[i]))
	}

	return input
}

func process(input []int) int {
	weight := 0
	available := make(map[int]bool, len(input))
	for _, w := range input {
		weight += w
		available[w] = true
	}

	for i := 1; i < len(input); i++ {
		products := find(i, weight/3, available, input)
		if len(products) == 0 {
			continue
		}
		min := input[0]
		for _, p := range products {
			min = aoc.Min(p)
		}
		return min
	}

	panic("Solution not found")
}

func find(parcelsLeft, requiredWeight int, available map[int]bool, parcels []int) []int {
	if parcelsLeft == 1 {
		if available[requiredWeight] {
			return []int{requiredWeight}
		}
		return nil
	}

	var result []int
	for _, w := range parcels {
		if w >= requiredWeight || !available[w] {
			continue
		}
		available[w] = false
		ap := find(parcelsLeft-1, requiredWeight-w, available, parcels)
		available[w] = true
		if len(ap) == 0 {
			continue
		}
		for _, p := range ap {
			result = append(result, p*w)
		}
	}

	return result
}
