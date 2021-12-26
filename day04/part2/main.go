package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() string {
	return aoc.ReadAllInput()[0]
}

func process(data string) int {
	for i := 0; ; i++ {
		p := []byte(fmt.Sprintf("%v%v", data, i))
		h := md5.Sum(p)
		s := fmt.Sprintf("%x", h)
		if strings.HasPrefix(s, "000000") {
			return i
		}
	}
}
