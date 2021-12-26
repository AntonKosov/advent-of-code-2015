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

func read() []aoc.Vector3 {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	var boxes []aoc.Vector3
	for _, b := range lines {
		sd := aoc.StrToInts(b, "x")
		boxes = append(boxes, aoc.NewVector3(sd[0], sd[1], sd[2]))
	}

	return boxes
}

func process(data []aoc.Vector3) int {
	sum := 0
	for _, box := range data {
		volume := box.X * box.Y * box.Z
		if box.X > box.Z {
			box.X, box.Z = box.Z, box.X
		}
		if box.Y > box.Z {
			box.Y, box.Z = box.Z, box.Y
		}

		length := 2 * (box.X + box.Y)
		sum += volume + length
	}

	return sum
}
