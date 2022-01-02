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

func read() (input int) {
	return aoc.StrToInt(aoc.ReadAllInput()[0])
}

func process(input int) int {
	const presents = 11
	const houses = 50
	for n := 1; n < input; n++ {
		sum := 0
		max := aoc.Sqrt(n)
		for i := 1; i <= max; i++ {
			if n%i != 0 {
				continue
			}
			if i <= houses {
				sum += n / i
			}
			if n/i <= houses {
				sum += i
			}
		}
		if sum*presents > input {
			return n
		}
	}

	panic("Solution not found")
}
