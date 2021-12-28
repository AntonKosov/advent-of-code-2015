package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (input map[int]map[int]int) {
	input = map[int]map[int]int{}

	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	cityIndexes := map[string]int{}
	cityIndex := func(city string) int {
		if i, ok := cityIndexes[city]; ok {
			return i
		}
		i := len(cityIndexes)
		cityIndexes[city] = i
		input[i] = make(map[int]int)
		return i
	}
	addDistance := func(city1, city2 string, distance int) {
		i1 := cityIndex(city1)
		i2 := cityIndex(city2)
		input[i1][i2] = distance
		input[i2][i1] = distance
	}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		addDistance(parts[0], parts[2], aoc.StrToInt(parts[4]))
	}

	return input
}

func process(input map[int]map[int]int) int {
	cities := make([]int, 0, len(input))
	for cID := range input {
		cities = append(cities, cID)
	}

	permutations := aoc.HeapPermutation(cities)
	max := 0
	for _, permutation := range permutations {
		dist := 0
		for i := 1; i < len(permutation); i++ {
			dist += input[permutation[i-1]][permutation[i]]
		}
		if dist > max {
			max = dist
		}
	}

	return max
}
