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
	m := make(map[aoc.Vector2]struct{}, len(data)*2)
	pos1 := aoc.NewVector2(0, 0)
	pos2 := pos1
	m[pos1] = struct{}{}

	for _, d := range data {
		switch d {
		case '<':
			pos1.X--
		case '>':
			pos1.X++
		case '^':
			pos1.Y--
		case 'v':
			pos1.Y++
		}
		m[pos1] = struct{}{}
		pos1, pos2 = pos2, pos1
	}

	return len(m)
}
