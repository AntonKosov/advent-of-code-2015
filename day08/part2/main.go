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

func read() (input []string) {
	lines := aoc.ReadAllInput()
	return lines[:len(lines)-1]
}

func process(input []string) int {
	diff := 0
	for _, word := range input {
		encoded := 2
		for _, r := range word {
			encoded++
			if r == '"' || r == '\\' {
				encoded++
			}
		}

		diff += encoded - len(word)
	}

	return diff
}
