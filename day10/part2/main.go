package main

import (
	"fmt"
	"strconv"

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
	runes := []rune(input)
	for i := 0; i < 50; i++ {
		next := make([]rune, 0, len(runes))
		for j := 0; j < len(runes); j++ {
			n := 1
			for j < len(runes)-1 && runes[j] == runes[j+1] {
				j++
				n++
			}
			value := []rune(strconv.Itoa(n))
			value = append(value, runes[j])
			next = append(next, value...)
		}

		runes = next
	}

	return len(runes)
}
