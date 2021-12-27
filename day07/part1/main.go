package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (input map[string][]string) { // target -> source
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	input = make(map[string][]string, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		target := parts[len(parts)-1]
		input[target] = parts[:len(parts)-2]
	}

	return input
}

func process(input map[string][]string) uint16 {
	variables := map[string]uint16{}
	return value("a", variables, input)
}

func value(st string, variables map[string]uint16, input map[string][]string) uint16 {
	if v, ok := variables[st]; ok {
		return v
	}
	if v, err := strconv.Atoi(st); err == nil {
		return uint16(v)
	}

	source := input[st]
	var v uint16

	switch len(source) {
	case 1:
		v = value(source[0], variables, input)
	case 2:
		v = ^value(source[1], variables, input)
	case 3:
		left := value(source[0], variables, input)
		right := value(source[2], variables, input)
		switch source[1] {
		case "OR":
			v = left | right
		case "AND":
			v = left & right
		case "RSHIFT":
			v = left >> right
		case "LSHIFT":
			v = left << right
		default:
			panic("Unknown command")
		}
	default:
		panic("Unknown command")
	}

	variables[st] = v

	return v
}
