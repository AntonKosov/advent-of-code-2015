package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (input [][]string) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		command := strings.Split(strings.ReplaceAll(line, ",", ""), " ")
		input = append(input, command)
	}

	return input
}

func process(input [][]string) int {
	regs := map[string]int{"a": 1, "b": 0}
	i := 0
	for i >= 0 && i < len(input) {
		command := input[i]
		arg1 := command[1]
		switch c := command[0]; c {
		case "hlf":
			regs[arg1] = regs[arg1] / 2
			i++
		case "tpl":
			regs[arg1] = regs[arg1] * 3
			i++
		case "inc":
			regs[arg1]++
			i++
		case "jmp":
			i += aoc.StrToInt(arg1)
		case "jie":
			if regs[arg1]%2 == 0 {
				i += aoc.StrToInt(command[2])
			} else {
				i++
			}
		case "jio":
			if regs[arg1] == 1 {
				i += aoc.StrToInt(command[2])
			} else {
				i++
			}
		default:
			panic("Unknown command")
		}
	}

	return regs["b"]
}
