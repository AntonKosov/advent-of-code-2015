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
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	badStart := map[rune]bool{'a': true, 'c': true, 'p': true, 'x': true}
	count := 0
	for _, word := range data {
		runes := []rune(word)
		v := 0
		d := 0
		n := false
		pr := runes[0]
		if vowels[pr] {
			v++
		}
		for i := 1; i < len(runes); i++ {
			cr := runes[i]
			if byte(cr-pr) == 1 && badStart[pr] {
				n = true
				break
			}
			if vowels[cr] {
				v++
			}
			if pr == cr {
				d++
			}

			pr = cr
		}
		if !n && v >= 3 && d > 0 {
			count++
		}
	}

	return count
}
