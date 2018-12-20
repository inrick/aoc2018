package main

import (
	"bufio"
	"fmt"
	"os"
)

type V2 struct{ x, y int }

var (
	North = V2{0, -1}
	East  = V2{1, 0}
	South = V2{0, 1}
	West  = V2{-1, 0}
)

func Add(u, v V2) V2 {
	return V2{u.x + v.x, u.y + v.y}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var input []byte
	if sc.Scan() {
		input = []byte(sc.Text())
	}

	pos := V2{0, 0}
	var re []V2 // Revisit stack
	edges := make(map[V2][]V2)
	for i := range input {
		switch input[i] {
		case 'N', 'E', 'S', 'W':
			from := pos
			switch input[i] {
			case 'N':
				pos = Add(pos, North)
			case 'E':
				pos = Add(pos, East)
			case 'S':
				pos = Add(pos, South)
			case 'W':
				pos = Add(pos, West)
			default:
				panic(nil)
			}
			edges[from] = append(edges[from], pos)
		case '(':
			re = append(re, pos)
		case ')':
			re = re[:len(re)-1]
		case '|':
			pos = re[len(re)-1]
		case '^':
			if i != 0 {
				panic(nil)
			}
		case '$':
			if i != len(input)-1 {
				panic(nil)
			}
		default:
			panic(nil)
		}
	}

	// BFS
	type Item struct {
		dist int
		node V2
	}
	Q := []Item{{0, V2{0, 0}}}
	seen := make(map[V2]bool)
	maxDist, thousandDoors := 0, 0
	for len(Q) > 0 {
		it := Q[0]
		Q = Q[1:]
		if it.dist > maxDist {
			maxDist = it.dist
		}
		if it.dist >= 1000 {
			thousandDoors++
		}
		for _, neighbor := range edges[it.node] {
			if !seen[neighbor] {
				Q = append(Q, Item{it.dist + 1, neighbor})
				seen[neighbor] = true
			}
		}
	}
	fmt.Printf("a) %d\n", maxDist)
	fmt.Printf("b) %d\n", thousandDoors)
}
