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

func read() []string {
	lines := aoc.ReadAllInput()
	return lines[:len(lines)-1]
}

func process(data []string) int {
	count := 0
	for _, word := range data {
		if isNice(word) {
			count++
		}
	}

	return count
}

func isNice(word string) bool {
	runes := []rune(word)
	repeats := false
	for i := 2; i < len(runes); i++ {
		if runes[i] == runes[i-2] {
			repeats = true
			break
		}
	}
	if !repeats {
		return false
	}
	for i := 1; i < len(runes)-2; i++ {
		r0 := runes[i-1]
		r1 := runes[i]
		for j := i + 2; j < len(runes); j++ {
			if r0 == runes[j-1] && r1 == runes[j] {
				return true
			}
		}
	}
	return false
}
