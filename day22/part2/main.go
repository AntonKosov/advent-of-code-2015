package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type player struct {
	hp        int
	mana      int
	armor     int
	spentMana int
}

func (p *player) spendMana(v int) {
	p.mana -= v
	p.spentMana += v
}

type boss struct {
	hp     int
	damage int
}

const (
	magicMissileManaCost = 53
	drainManaCost        = 73
	shieldManaCost       = 113
	poisonManaCost       = 173
	rechargeManaCost     = 229
)

type action func(*player, *boss)

type effect struct {
	left   int
	action action
}

const maxEffects = 3

type effectID int

const (
	noneEffectID     = -1
	shieldEffectID   = 0
	poisonEffectID   = 1
	rechargeEffectID = 2
)

type effects struct {
	list [maxEffects]effect
}

func (e *effects) apply(p *player, b *boss) {
	p.armor = 0
	for i := 0; i < len(e.list); i++ {
		eff := &e.list[i]
		if eff.left == 0 {
			continue
		}
		eff.action(p, b)
		eff.left--
	}
}

func shieldEffect(p *player, _ *boss) {
	p.armor = 7
}

func poisonEffect(_ *player, b *boss) {
	b.hp -= 3
}

func rechargeEffect(p *player, _ *boss) {
	p.mana += 101
}

func (e *effects) add(id effectID) {
	var newEffect effect
	switch id {
	case shieldEffectID:
		newEffect = effect{left: 6, action: shieldEffect}
	case poisonEffectID:
		newEffect = effect{left: 6, action: poisonEffect}
	case rechargeEffectID:
		newEffect = effect{left: 5, action: rechargeEffect}
	default:
		panic("Unknown effect")
	}

	e.list[id] = newEffect
}

func (e *effects) isActive(id effectID) bool {
	return e.list[id].left > 0
}

type state struct {
	player  player
	boss    boss
	effects effects
}

func (s state) isBossDead() bool {
	return s.boss.hp <= 0
}

func (s state) isPlayerDead() bool {
	return s.player.hp <= 0
}

type states []state

func read() boss {
	lines := aoc.ReadAllInput()
	p := func(line string) int {
		parts := strings.Split(line, " ")
		return aoc.StrToInt(parts[len(parts)-1])
	}
	return boss{
		hp:     p(lines[0]),
		damage: p(lines[1]),
	}
}

func process(boss boss) int {
	min := math.MaxInt
	currentStates := states{{
		player:  player{hp: 50, mana: 500},
		boss:    boss,
		effects: effects{},
	}}

	iterations := 0
	for len(currentStates) > 0 {
		nextStates := make(states, 0, len(currentStates))
		for _, s := range currentStates {
			if s.player.spentMana >= min {
				continue
			}
			s.player.hp--
			if s.isPlayerDead() {
				continue
			}
			s.effects.apply(&s.player, &s.boss)
			if s.isBossDead() {
				min = aoc.Min(min, s.player.spentMana)
				continue
			}
			newStates := make(states, 0, 5)
			addNewState := func(ns state, manaCost, playerHPChange, bossHPChange int, e effectID) {
				if ns.player.mana < manaCost {
					return
				}
				if e != noneEffectID {
					if ns.effects.isActive(e) {
						return
					}
					ns.effects.add(e)
				}
				ns.player.spendMana(manaCost)
				ns.player.hp += playerHPChange
				ns.boss.hp += bossHPChange
				ns.effects.apply(&ns.player, &ns.boss)
				if ns.isBossDead() {
					min = aoc.Min(min, ns.player.spentMana)
					return
				}
				bossDamage := aoc.Max(1, ns.boss.damage-ns.player.armor)
				ns.player.hp -= bossDamage
				if ns.isPlayerDead() {
					return
				}
				newStates = append(newStates, ns)
			}
			addNewState(s, magicMissileManaCost, 0, -4, noneEffectID)
			addNewState(s, drainManaCost, 2, -2, noneEffectID)
			addNewState(s, shieldManaCost, 0, 0, shieldEffectID)
			addNewState(s, poisonManaCost, 0, 0, poisonEffectID)
			addNewState(s, rechargeManaCost, 0, 0, rechargeEffectID)

			nextStates = append(nextStates, newStates...)
		}

		currentStates = nextStates
		iterations++
	}

	return min
}
