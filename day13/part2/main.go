package main

import (
	"fmt"
	"math"
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

	peopleIndexes := map[string]int{}
	personIndex := func(name string) int {
		if i, ok := peopleIndexes[name]; ok {
			return i
		}
		i := len(peopleIndexes)
		peopleIndexes[name] = i
		input[i] = make(map[int]int)
		return i
	}
	addHappines := func(nameFrom, nameTo string, happines int) {
		i1 := personIndex(nameFrom)
		i2 := personIndex(nameTo)
		input[i1][i2] = happines
	}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		from := parts[0]
		to := parts[len(parts)-1]
		to = to[:len(to)-1]
		happines := aoc.StrToInt(parts[3])
		if parts[2] == "lose" {
			happines = -happines
		}
		addHappines(from, to, happines)
	}
	personIndex("me")

	return input
}

func process(input map[int]map[int]int) int {
	people := make([]int, 0, len(input))
	for cID := range input {
		people = append(people, cID)
	}

	permutations := aoc.HeapPermutation(people)
	max := math.MinInt
	for _, permutation := range permutations {
		happines := 0
		for i := 0; i < len(permutation); i++ {
			p1 := permutation[i]
			p2 := permutation[(i+1)%len(people)]
			happines += input[p1][p2]
			happines += input[p2][p1]
		}
		max = aoc.Max(max, happines)
	}

	return max
}
