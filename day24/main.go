package main

import (
	"fmt"
	"sort"
)

type System int

const (
	Immune System = iota
	Infection
)

type Attack int

const (
	Cold Attack = iota
	Fire
	Radiation
	Bludgeoning
	Slashing
)

type Group struct {
	sys    System
	units  int
	hp     int
	immune []Attack
	weak   []Attack
	attack Attack
	damage int
	init   int
}

func main() {
	// Took the lazy route and parsed the input manually. Some editor search &
	// replace fixed most of it.
	groups := []Group{
		{Immune, 933, 3691, nil, nil, Cold, 37, 15},
		{Immune, 262, 2029, nil, nil, Cold, 77, 4},
		{Immune, 3108, 2902, []Attack{Slashing, Fire}, []Attack{Bludgeoning}, Cold, 7, 13},
		{Immune, 5158, 9372, []Attack{Cold}, []Attack{Bludgeoning}, Radiation, 17, 16},
		{Immune, 2856, 4797, nil, nil, Cold, 16, 20},
		{Immune, 86, 8311, nil, nil, Slashing, 724, 14},
		{Immune, 7800, 3616, []Attack{Radiation, Cold, Bludgeoning}, nil, Bludgeoning, 4, 7},
		{Immune, 1374, 8628, nil, []Attack{Fire, Slashing}, Radiation, 61, 1},
		{Immune, 1661, 4723, nil, nil, Slashing, 25, 8},
		{Immune, 1285, 4156, nil, []Attack{Bludgeoning}, Fire, 32, 18},
		{Infection, 2618, 29001, []Attack{Bludgeoning, Radiation, Cold}, nil, Radiation, 17, 3},
		{Infection, 31, 20064, []Attack{Slashing, Bludgeoning}, []Attack{Radiation}, Bludgeoning, 1082, 10},
		{Infection, 281, 15311, nil, []Attack{Fire, Cold}, Slashing, 90, 9},
		{Infection, 1087, 14744, []Attack{Radiation}, []Attack{Cold, Fire}, Fire, 22, 12},
		{Infection, 7810, 48137, nil, []Attack{Fire, Radiation}, Slashing, 10, 5},
		{Infection, 232, 18762, nil, []Attack{Radiation, Cold}, Bludgeoning, 153, 2},
		{Infection, 69, 11032, []Attack{Radiation, Slashing, Cold, Fire}, nil, Slashing, 296, 6},
		{Infection, 2993, 10747, []Attack{Slashing}, []Attack{Cold}, Radiation, 6, 19},
		{Infection, 273, 7590, []Attack{Slashing, Fire}, []Attack{Radiation}, Fire, 49, 17},
		{Infection, 2041, 38432, nil, []Attack{Bludgeoning}, Cold, 29, 11},
	}

	// For part b. Binary search by hand. Note that it is possible to reach a
	// steady state where neither the infection or immune system wins and it just
	// keeps looping forever.
	/*
		boost := 33
		for i, _ := range groups {
			if groups[i].sys == Immune {
				groups[i].damage += boost
			}
		}
	*/

	order := make([]int, len(groups))
	for i := range order {
		order[i] = i
	}

	for {
		g1, g2 := false, false
		for _, g := range groups {
			if g.sys == Immune && g.units > 0 {
				g1 = true
			} else if g.sys == Infection && g.units > 0 {
				g2 = true
			}
		}
		if !g1 || !g2 {
			break
		}

		// Target selection
		sort.Slice(order, func(i, j int) bool {
			i, j = order[i], order[j]
			ei, ej := Eff(groups[i]), Eff(groups[j])
			return ei > ej || (ei == ej && groups[i].init > groups[j].init)
		})

		// Only save multiplier, rather than damage, since the attacker can lose
		// units before the attack takes place, thus losing effective power.
		type Target struct {
			j, mult int
		}
		targets := make([]Target, len(groups))
		marked := make([]bool, len(groups))

		for _, i := range order {
			a := &groups[i]
			if a.units <= 0 {
				continue
			}
			maxdmg := -1
			target := Target{-1, -1}
			for j := range groups {
				b := &groups[j]
				// Skip if same type or already marked for attack or dead
				if a.sys == b.sys || marked[j] || b.units <= 0 {
					continue
				}
				mult := 1
				if contains(a.attack, b.immune) {
					mult = 0
				} else if contains(a.attack, b.weak) {
					mult = 2
				}
				dmg := mult * Eff(*a)
				if 0 < dmg && maxdmg < dmg || (maxdmg == dmg &&
					(Eff(groups[target.j]) < Eff(*b) ||
						(Eff(groups[target.j]) == Eff(*b) && groups[target.j].init < b.init))) {
					target = Target{j, mult}
					maxdmg = dmg
				}
			}
			targets[i] = target
			if target.j != -1 {
				marked[target.j] = true
			}
		}

		// Attack phase
		sort.Slice(order, func(i, j int) bool {
			return groups[order[i]].init > groups[order[j]].init
		})
		for _, i := range order {
			if groups[i].units <= 0 || targets[i].j == -1 {
				continue
			}
			target := &groups[targets[i].j]
			killed := int(targets[i].mult * Eff(groups[i]) / target.hp)
			target.units -= killed
		}
	}

	sum := 0
	for _, g := range groups {
		if g.units > 0 {
			/*if g.sys != Immune {
				panic(nil)
			}*/
			sum += g.units
		}
	}
	fmt.Printf("a) %d\n", sum)
}

func contains(a Attack, as []Attack) bool {
	for _, b := range as {
		if a == b {
			return true
		}
	}
	return false
}

func Eff(g Group) int {
	return g.units * g.damage
}
