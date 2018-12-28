package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type V3 struct {
	x, y, z int
}

type Bot struct {
	pos V3
	r   int
}

var Origin = V3{0, 0, 0}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var bots []Bot
	for sc.Scan() {
		var b Bot
		fmt.Sscanf(sc.Text(), "pos=<%d,%d,%d>, r=%d", &b.pos.x, &b.pos.y, &b.pos.z, &b.r)
		bots = append(bots, b)
	}

	// Reverse sort
	order := make([]int, len(bots))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		return bots[order[i]].r > bots[order[j]].r
	})
	// a
	strongest := bots[order[0]]
	sum := 0
	for _, b := range bots {
		if Dist(strongest.pos, b.pos) <= strongest.r {
			sum++
		}
	}
	fmt.Printf("a) %d\n", sum)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Dist(u, v V3) int {
	return Abs(u.x-v.x) + Abs(u.y-v.y) + Abs(u.z-v.z)
}
