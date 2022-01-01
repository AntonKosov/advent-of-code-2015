package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type formula []string

type formulas struct {
	conversions map[string][]formula
	formula     formula
}

func read() (input formulas) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	input.conversions = map[string][]formula{}
	conversionSplitter := regexp.MustCompile(`\w+`)
	formulaSplitter := regexp.MustCompile(`[A-Z][a-z]?`)

	for _, line := range lines {
		if line == "" {
			break
		}
		conversionParts := conversionSplitter.FindAllString(line, -1)
		from := conversionParts[0]
		to := formulaSplitter.FindAllString(conversionParts[1], -1)
		input.conversions[from] = append(input.conversions[from], to)
	}

	input.formula = formulaSplitter.FindAllString(lines[len(lines)-1], -1)

	return input
}

func process(input formulas) int {
	variants := map[string]struct{}{}
	for i, ch := range input.formula {
		replacement, ok := input.conversions[ch]
		if !ok {
			continue
		}
		firstPart := ""
		if i > 0 {
			firstPart = strings.Join(input.formula[:i], "")
		}
		lastPart := ""
		if i < len(input.formula)-1 {
			lastPart = strings.Join(input.formula[i+1:], "")
		}
		for _, res := range replacement {
			middlePart := strings.Join(res, "")
			full := fmt.Sprintf("%v%v%v", firstPart, middlePart, lastPart)
			variants[full] = struct{}{}
		}
	}

	return len(variants)
}
