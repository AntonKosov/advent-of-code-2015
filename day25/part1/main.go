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

func read() (input aoc.Vector2) {
	line := aoc.ReadAllInput()[0]
	rep := regexp.MustCompile(`\d+`)
	parts := rep.FindAllString(line, -1)
	row := aoc.StrToInt(parts[0])
	col := aoc.StrToInt(parts[1])

	return aoc.NewVector2(col, row)
}

func process(input aoc.Vector2) int {
	value := 20151125
	for row := 2; ; row++ {
		for col := 1; col <= row; col++ {
			value = (value * 252533) % 33554393
			if input.X == col && input.Y == row-col+1 {
				return value
			}
		}
	}
}
