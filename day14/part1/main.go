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

type reindeer struct {
	speed      int
	flyingTime int
	restTime   int
}

func read() (input []reindeer) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		parts := strings.Split(line, " ")
		input = append(input, reindeer{
			speed:      aoc.StrToInt(parts[3]),
			flyingTime: aoc.StrToInt(parts[6]),
			restTime:   aoc.StrToInt(parts[13]),
		})
	}

	return input
}

func process(input []reindeer) int {
	const timeLimit = 2503
	distances := make([]int, len(input))
	for i, reindeer := range input {
		distance := 0
		for time := 0; time < timeLimit; {
			if time > 0 {
				time += reindeer.restTime
				if time >= timeLimit {
					break
				}
			}
			flyingTime := aoc.Min(timeLimit-time, reindeer.flyingTime)
			distance += reindeer.speed * flyingTime
			time += flyingTime
		}
		distances[i] = distance
	}

	max := aoc.Max(distances...)

	return max
}
