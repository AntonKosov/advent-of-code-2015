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

func read() []string {
	lines := aoc.ReadAllInput()
	formula := lines[len(lines)-2]

	return parseFormula(formula)
}

func parseFormula(formula string) []string {
	formulaSplitter := regexp.MustCompile(`[A-Z][a-z]?`)

	return formulaSplitter.FindAllString(formula, -1)
}

const (
	start     = "Rn"
	end       = "Ar"
	delimiter = "Y"
)

func process(in []string) int {
	steps := 0
	for i := 0; i < len(in)-1; {
		if in[i+1] != start {
			i++
			steps++
			continue
		}
		s, lastIndex := countBrackets(in, i)
		steps += s
		i = lastIndex
	}

	return steps
}

func countBrackets(in []string, startIndex int) (steps, lastIndex int) {
	i := startIndex + 2
	for {
		switch next := in[i+1]; next {
		case start:
			s, li := countBrackets(in, i)
			steps += s
			i = li
		case end:
			return steps + 1, i + 1
		case delimiter:
			i += 2
		default:
			i++
			steps++
		}
	}
}
