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

type things map[string]int // thing -> count

func read() (input map[int]things) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	input = make(map[int]things, len(lines))
	re := regexp.MustCompile(`(\d|\w)+`)

	for _, line := range lines {
		parts := re.FindAll([]byte(line), -1)
		t := things{}
		for i := 2; i < len(parts); i += 2 {
			thing := string(parts[i])
			count := aoc.StrToInt(string(parts[i+1]))
			t[thing] = count
		}

		input[aoc.StrToInt(string(parts[1]))] = t
	}

	return input
}

func process(input map[int]things) int {
	known := things{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

nextAunt:
	for id, t := range input {
		for name, count := range t {
			if known[name] != count {
				continue nextAunt
			}
		}
		return id
	}

	panic("No solution")
}
