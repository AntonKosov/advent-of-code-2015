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

type character struct {
	damage int
	armor  int
	hp     int
}

func read() (boss character) {
	lines := aoc.ReadAllInput()
	p := func(line string) int {
		parts := strings.Split(line, " ")
		return aoc.StrToInt(parts[len(parts)-1])
	}
	return character{
		hp:     p(lines[0]),
		damage: p(lines[1]),
		armor:  p(lines[2]),
	}
}

type item struct {
	price  int
	damage int
	armor  int
}

var (
	weapons []item
	armors  []item
	rings   []item
)

func init() {
	weapons = []item{
		{price: 8, damage: 4},
		{price: 10, damage: 5},
		{price: 25, damage: 6},
		{price: 40, damage: 7},
		{price: 74, damage: 8},
	}
	armors = []item{
		{},
		{price: 13, armor: 1},
		{price: 31, armor: 2},
		{price: 53, armor: 3},
		{price: 75, armor: 4},
		{price: 102, armor: 5},
	}
	rings = []item{
		{},
		{},
		{price: 25, damage: 1},
		{price: 50, damage: 2},
		{price: 100, damage: 3},
		{price: 20, armor: 1},
		{price: 40, armor: 2},
		{price: 80, armor: 3},
	}

}

func process(boss character) int {
	max := 0

	for _, weapon := range weapons {
		for _, armor := range armors {
			for r1 := 0; r1 < len(rings)-1; r1++ {
				ring1 := rings[r1]
				for r2 := r1 + 1; r2 < len(rings); r2++ {
					ring2 := rings[r2]
					player := character{
						hp:     100,
						damage: weapon.damage + ring1.damage + ring2.damage,
						armor:  armor.armor + ring1.armor + ring2.armor,
					}
					if !isWinner(player, boss) {
						cost := weapon.price + armor.price + ring1.price + ring2.price
						max = aoc.Max(max, cost)
					}
				}
			}
		}
	}

	return max
}

func isWinner(player, boss character) bool {
	attack := func(attacker character, defender *character) bool {
		damage := aoc.Max(1, attacker.damage-defender.armor)
		defender.hp -= damage
		return defender.hp <= 0
	}

	for {
		if attack(player, &boss) {
			return true
		}
		if attack(boss, &player) {
			return false
		}
	}
}
