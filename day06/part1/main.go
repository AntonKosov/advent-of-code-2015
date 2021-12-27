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

type command int

const (
	commandToggle  command = 0
	commandTurnOn  command = 1
	commandTurnOff command = 2
)

type instruction struct {
	command command
	min     aoc.Vector2
	max     aoc.Vector2
}

func read() (input []instruction) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	input = make([]instruction, 0, len(lines))
	for _, inst := range lines {
		i := instruction{}
		parts := strings.Split(inst, " ")
		if len(parts) == 4 {
			i.command = commandToggle
		} else if parts[1] == "on" {
			i.command = commandTurnOn
		} else {
			i.command = commandTurnOff
		}

		min := aoc.StrToInts(parts[len(parts)-3], ",")
		max := aoc.StrToInts(parts[len(parts)-1], ",")
		i.min = aoc.NewVector2(min[0], min[1])
		i.max = aoc.NewVector2(max[0], max[1])

		input = append(input, i)
	}

	return input
}

func process(input []instruction) int {
	m := [1000][1000]bool{}
	for _, inst := range input {
		for x := inst.min.X; x <= inst.max.X; x++ {
			for y := inst.min.Y; y <= inst.max.Y; y++ {
				switch inst.command {
				case commandToggle:
					m[x][y] = !m[x][y]
				case commandTurnOn:
					m[x][y] = true
				case commandTurnOff:
					m[x][y] = false
				default:
					panic("Unknown command")
				}
			}
		}
	}

	sum := 0
	for _, col := range m {
		for _, v := range col {
			if v {
				sum++
			}
		}
	}

	return sum
}
