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

func read() (input string) {
	return aoc.ReadAllInput()[0]
}

func process(input string) string {
	psw := []rune(input)

	for i, r := range psw {
		if !skip[r] {
			continue
		}
		psw[i]++
		resetToA(psw, i+1)
		break
	}

	for !(hasIncreasingLetters(psw) && hasPairs(psw)) {
		inc(psw)
	}

	return string(psw)
}

var skip map[rune]bool

func init() {
	skip = map[rune]bool{
		'i': true,
		'o': true,
		'l': true,
	}
}

func resetToA(psw []rune, start int) {
	for i := start; i < len(psw); i++ {
		psw[i] = 'a'
	}
}

func inc(psw []rune) {
	for i := len(psw) - 1; i >= 0; i-- {
		nl := psw[i] + 1
		if nl > 'z' {
			psw[i] = 'a'
			continue
		}
		if skip[nl] {
			nl++
			resetToA(psw, i+1)
		}
		psw[i] = nl
		break
	}
}

func hasIncreasingLetters(psw []rune) bool {
	for i := len(psw) - 3; i >= 0; i-- {
		if psw[i+2]-psw[i+1] == 1 && psw[i+1]-psw[i] == 1 {
			return true
		}
	}

	return false
}

func hasPairs(psw []rune) bool {
	for i := 0; i <= len(psw)-4; i++ {
		if psw[i] != psw[i+1] {
			continue
		}
		for j := i + 2; j < len(psw)-1; j++ {
			if psw[j] == psw[j+1] {
				return true
			}
		}
	}

	return false
}
