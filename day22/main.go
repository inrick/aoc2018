package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type V2 struct {
	x, y int
}

func Add(u, v V2) V2 {
	return V2{u.x + v.x, u.y + v.y}
}

var (
	Up    = V2{0, -1}
	Right = V2{1, 0}
	Left  = V2{-1, 0}
	Down  = V2{0, 1}
)

type Gear int

const (
	Neither Gear = iota
	Climbing
	Torch
)

const (
	Rocky  = 0
	Wet    = 1
	Narrow = 2
)

func main() {
	var depth int
	var target V2
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fmt.Sscanf(sc.Text(), "depth: %d", &depth)
	sc.Scan()
	fmt.Sscanf(sc.Text(), "target: %d,%d", &target.x, &target.y)

	mem := make(map[V2]int)
	regType := func(u V2) int {
		return Erosion(mem, depth, target, u) % 3
	}

	risk := 0
	for y := 0; y <= target.y; y++ {
		for x := 0; x <= target.x; x++ {
			risk += regType(V2{x, y})
		}
	}
	fmt.Printf("a) %d\n", risk)

	// b

	// Need both position and gear when bfs:ing as the gear choice cascades.
	type Key struct {
		gear Gear
		u    V2
	}
	type Item struct {
		dist int
		typ  int
		Key
	}

	dist := 1 << 30
	visited := make(map[Key]bool)
	Q := []Item{Item{0, 0, Key{Torch, V2{0, 0}}}}
	var it Item
	for 0 < len(Q) {
		sort.Slice(Q, func(i, j int) bool { return Q[i].dist < Q[j].dist })
		it, Q = Q[0], Q[1:]
		for visited[it.Key] {
			it, Q = Q[0], Q[1:]
		}
		visited[it.Key] = true
		if it.u == target {
			dist = it.dist
			if it.gear != Torch {
				dist += 7
				// Might get here quicker through another route with proper gear
				// equipped
				continue
			}
			break
		}
		if it.dist > dist {
			// Nah, already found shortest route
			break
		}

		for _, neighbor := range Neighbors(it.u) {
			// Rather ugly, drop paths that stray according to made up numbers...
			if neighbor.x < 0 || neighbor.y < 0 ||
				neighbor.x > target.x+100 || neighbor.y > target.y+1000 {
				continue
			}
			ntyp := regType(neighbor)
			var dt int
			var key Key
			switch {
			case it.typ == ntyp:
				dt = 1
				key = Key{it.gear, neighbor}
			case (it.typ == Rocky && ntyp == Wet) || (it.typ == Wet && ntyp == Rocky):
				dt = 1
				if it.gear != Climbing {
					dt = 8
				}
				key = Key{Climbing, neighbor}
			case (it.typ == Rocky && ntyp == Narrow) || (it.typ == Narrow && ntyp == Rocky):
				dt = 1
				if it.gear != Torch {
					dt = 8
				}
				key = Key{Torch, neighbor}
			case (it.typ == Wet && ntyp == Narrow) || (it.typ == Narrow && ntyp == Wet):
				dt = 1
				if it.gear != Neither {
					dt = 8
				}
				key = Key{Neither, neighbor}
			default:
				// Nothing to do
			}
			if dt > 0 && !visited[key] {
				Q = append(Q, Item{dt + it.dist, ntyp, key})
			}
		}
	}

	fmt.Printf("b) %d\n", dist)
}

func GeoIndex(mem map[V2]int, depth int, target, at V2) int {
	m, ok := mem[at]
	if ok {
		return m
	}
	var index int
	switch {
	case at == V2{0, 0} || at == target:
		index = 0
	case at.y == 0:
		index = at.x * 16807
	case at.x == 0:
		index = at.y * 48271
	default:
		index = Erosion(mem, depth, target, V2{at.x - 1, at.y}) *
			Erosion(mem, depth, target, V2{at.x, at.y - 1})
	}
	mem[at] = index
	return index
}

func Erosion(mem map[V2]int, depth int, target, at V2) int {
	return (GeoIndex(mem, depth, target, at) + depth) % 20183
}

func Neighbors(u V2) []V2 {
	return []V2{Add(u, Up), Add(u, Left), Add(u, Down), Add(u, Right)}
}
