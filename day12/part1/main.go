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

func read() (input string) {
	return aoc.ReadAllInput()[0]
}

func process(input string) int {
	m := regexp.MustCompile(`-?\d+`)
	values := m.FindAllString(input, -1)
	sum := 0
	for _, v := range values {
		sum += aoc.StrToInt(v)
	}
	return sum
}
