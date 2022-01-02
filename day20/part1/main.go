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

const presents = 10

func process(input int) int {
	for n := 10; ; n++ {
		sum := 0
		max := aoc.Sqrt(n)
		for i := 1; i <= max; i++ {
			if n%i != 0 {
				continue
			}
			sum += i
			s := n / i
			if s != i {
				sum += s
			}
		}
		if sum*presents > input {
			return n
		}
	}
}
