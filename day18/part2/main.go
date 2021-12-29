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

type field struct {
	lit    map[aoc.Vector2]int
	width  int
	height int
}

func read() (input field) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	input.height = len(lines)
	input.width = len(lines[0])
	input.lit = map[aoc.Vector2]int{}

	for y, line := range lines {
		for x, v := range line {
			if v == '#' {
				input.lit[aoc.NewVector2(x, y)] = 1
			}
		}
	}

	return input
}

func process(input field) int {
	alwaysOn := map[aoc.Vector2]bool{
		aoc.NewVector2(0, 0):                          true,
		aoc.NewVector2(0, input.height-1):             true,
		aoc.NewVector2(input.width-1, 0):              true,
		aoc.NewVector2(input.width-1, input.height-1): true,
	}
	for i := 0; i < 100; i++ {
		var on []aoc.Vector2
		var off []aoc.Vector2
		for x := 0; x < input.width; x++ {
			for y := 0; y < input.height; y++ {
				n := 0
				p := aoc.NewVector2(x, y)
				for _, a := range p.Adjacent() {
					n += input.lit[a]
				}
				isOn := input.lit[p] == 1
				if n == 3 && !isOn {
					on = append(on, p)
				} else if n != 2 && n != 3 && isOn && !alwaysOn[p] {
					off = append(off, p)
				}
			}
		}
		for _, p := range off {
			delete(input.lit, p)
		}
		for _, p := range on {
			input.lit[p] = 1
		}
	}

	return len(input.lit)
}
