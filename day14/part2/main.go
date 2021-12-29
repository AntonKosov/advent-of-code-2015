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

	flyingTimeLeft int
	restTimeLeft   int
	distance       int
	points         int
}

func (r *reindeer) turn() {
	if r.restTimeLeft > 0 {
		r.restTimeLeft--
		if r.restTimeLeft == 0 {
			r.flyingTimeLeft = r.flyingTime
		}
		return
	}

	r.flyingTimeLeft--
	if r.flyingTimeLeft == 0 {
		r.restTimeLeft = r.restTime
	}

	r.distance += r.speed
}

func read() (input []*reindeer) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		parts := strings.Split(line, " ")
		r := &reindeer{
			speed:      aoc.StrToInt(parts[3]),
			flyingTime: aoc.StrToInt(parts[6]),
			restTime:   aoc.StrToInt(parts[13]),
		}
		r.flyingTimeLeft = r.flyingTime
		input = append(input, r)
	}

	return input
}

func process(reindeers []*reindeer) int {
	const timeLimit = 2503
	for t := 0; t < timeLimit; t++ {
		maxDistance := 0
		for _, r := range reindeers {
			r.turn()
			maxDistance = aoc.Max(maxDistance, r.distance)
		}
		for _, r := range reindeers {
			if r.distance == maxDistance {
				r.points++
			}
		}
	}

	max := 0
	for _, r := range reindeers {
		max = aoc.Max(max, r.points)
	}

	return max
}
