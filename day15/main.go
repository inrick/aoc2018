package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const Debug = false

type Type int

const (
	Elf Type = iota
	Goblin
)

type Pt struct {
	x, y int
}

type Unit struct {
	ty  Type
	pos Pt
	ap  int // attack power
	hp  int // hit points
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var cave [][]byte
	for sc.Scan() {
		cave = append(cave, []byte(sc.Text()))
	}

	// # = wall
	// . = open cavern
	// E = elf
	// G = goblin

	var units []Unit
	for y := range cave {
		for x := range cave[y] {
			switch cave[y][x] {
			case 'E', 'G':
				var ty Type
				switch cave[y][x] {
				case 'E':
					ty = Elf
				case 'G':
					ty = Goblin
				}
				cave[y][x] = '.'
				units = append(units, Unit{ty, Pt{x, y}, 3, 200})
			}
		}
	}

	// Reverse lookup: point -> unit
	at := make([][]int, len(cave))
	for y := range cave {
		at[y] = make([]int, len(cave[y]))
		for x := range cave[y] {
			at[y][x] = -1
		}
	}
	for i, u := range units {
		at[u.pos.y][u.pos.x] = i
	}

	order := make([]int, len(units))
	for i := range order {
		order[i] = i
	}

	finalRound := -1

Game:
	for round := 0; ; round++ {
		// Sort in reading order
		sort.Slice(order, func(i, j int) bool {
			return IsBefore(units[order[i]].pos, units[order[j]].pos)
		})

		if Debug {
			fmt.Println("After", round, "rounds:")
			PrintCave(cave, units)
		}

	Round:
		for _, i := range order {
			u := &units[i]
			if u.hp <= 0 {
				continue
			}

			enemiesLeft := false
			for _, v := range units {
				if v.hp > 0 && v.ty != u.ty {
					enemiesLeft = true
					break
				}
			}
			if !enemiesLeft {
				finalRound = round
				break Game
			}

			if Attack(u, units, at) {
				continue Round
			}

			// No adjacent enemy, try to make a move.
			previous := make(map[Pt]Pt)
			visited := make(map[Pt]bool)
			Q := make([]Pt, 1)
			Q[0] = u.pos

			var paths [][]Pt

			for 0 < len(Q) {
				visit := Q[0]
				Q = Q[1:]
				if visited[visit] {
					continue
				}
				visited[visit] = true

				for _, adjacent := range Adjacent(visit) {
					if !visited[adjacent] && cave[adjacent.y][adjacent.x] == '.' {
						if v := at[adjacent.y][adjacent.x]; v != -1 {
							if units[v].ty != u.ty {
								// Found an enemy, search backward to find path.
								path := []Pt{visit}
								for prev := visit; previous[prev] != u.pos; {
									prev = previous[prev]
									path = append(path, prev)
								}
								// Note that path is reversed
								paths = append(paths, path)
							}
						} else if _, ok := previous[adjacent]; !ok {
							// We are not on track to visit this this point yet
							previous[adjacent] = visit
							Q = append(Q, adjacent)
						}
						// Else further travel is blocked off
					}
				}
			}

			if len(paths) > 0 {
				// Find shortest path where target is first in reading order
				sort.Slice(paths, func(i, j int) bool {
					Ni, Nj := len(paths[i]), len(paths[j])
					return Ni < Nj || (Ni == Nj && IsBefore(paths[i][0], paths[j][0]))
				})

				moveTo := paths[0][len(paths[0])-1]
				at[u.pos.y][u.pos.x] = -1
				u.pos = moveTo
				at[moveTo.y][moveTo.x] = i

				// Try to attack after moving unit
				Attack(u, units, at)
			}
		}
	}

	if Debug {
		fmt.Println("Game finished in round", finalRound, "with state:")
		PrintCave(cave, units)
	}

	hp := 0
	for _, u := range units {
		if u.hp > 0 {
			hp += u.hp
		}
	}
	fmt.Printf("a) %d\n", finalRound*hp)
}

// Return adjacent points in reading order
func Adjacent(pt Pt) []Pt {
	x, y := pt.x, pt.y
	return []Pt{{x, y - 1}, {x - 1, y}, {x + 1, y}, {x, y + 1}}
}

func IsBefore(p, q Pt) bool {
	return p.y < q.y || (p.y == q.y && p.x < q.x)
}

func Attack(u *Unit, units []Unit, at [][]int) bool {
	// See if an enemy is adjacent
	var enemies []int
	for _, pt := range Adjacent(u.pos) {
		if j := at[pt.y][pt.x]; j != -1 {
			if u.ty != units[j].ty {
				enemies = append(enemies, j)
			}
		}
	}
	if len(enemies) > 0 {
		// Attack enemy with lowest hp, or closest
		sort.Slice(enemies, func(a, b int) bool {
			e0 := units[enemies[a]]
			e1 := units[enemies[b]]
			return e0.hp < e1.hp || (e0.hp == e1.hp && IsBefore(e0.pos, e1.pos))
		})
		v := &units[enemies[0]]
		v.hp -= u.ap
		if v.hp <= 0 {
			// Dead
			at[v.pos.y][v.pos.x] = -1
			v.pos = Pt{-1, -1}
		}
		return true
	}
	return false
}

func PrintCave(cave [][]byte, units []Unit) {
	output := make([][]byte, len(cave))
	for i := range output {
		output[i] = make([]byte, len(cave[i]))
		copy(output[i], cave[i])
	}
	for _, u := range units {
		if u.hp > 0 {
			output[u.pos.y][u.pos.x] = Output(u.ty)
		}
	}
	for _, line := range output {
		fmt.Println(string(line))
	}
}

func Output(u Type) byte {
	switch u {
	case Elf:
		return 'E'
	case Goblin:
		return 'G'
	default:
		panic(nil)
	}
}
