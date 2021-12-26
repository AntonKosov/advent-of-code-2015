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

func read() string {
	return aoc.ReadAllInput()[0]
}

func process(data string) int {
	m := make(map[aoc.Vector2]struct{}, len(data)+1)
	pos := aoc.NewVector2(0, 0)
	m[pos] = struct{}{}

	for _, d := range data {
		switch d {
		case '<':
			pos.X--
		case '>':
			pos.X++
		case '^':
			pos.Y--
		case 'v':
			pos.Y++
		}
		m[pos] = struct{}{}
	}

	return len(m)
}
