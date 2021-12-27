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
		mem := 0
		runes := []rune(word)
		for i := 1; i < len(runes)-1; i++ {
			mem++
			c := runes[i]
			if c != '\\' {
				continue
			}
			switch n := runes[i+1]; n {
			case '\\':
				i++
			case '"':
				i++
			case 'x':
				i += 3
			}
		}

		diff += len(word) - mem
	}

	return diff
}
