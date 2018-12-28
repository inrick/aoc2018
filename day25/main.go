package main

import (
	"bufio"
	"fmt"
	"os"
)

type v4 struct {
	x, y, z, w int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var input []v4
	for sc.Scan() {
		var u v4
		fmt.Sscanf(sc.Text(), "%d,%d,%d,%d", &u.x, &u.y, &u.z, &u.w)
		input = append(input, u)
	}

	// Build graph
	edges := make([][]int, len(input))
	for i := range input {
	NeighborLook:
		for j := range input {
			if i == j {
				continue
			}
			if dist(input[i], input[j]) <= 3 {
				for _, k := range edges[i] {
					if k == j {
						continue NeighborLook
					}
				}
				edges[i] = append(edges[i], j)
			}
		}
	}

	visited := make([]bool, len(input))
	constellations := 0
	for i := range visited {
		if visited[i] {
			continue
		}
		Q := []int{i}
		for 0 < len(Q) {
			visit := Q[0]
			Q = Q[1:]
			visited[visit] = true
			for _, neighbor := range edges[visit] {
				if !visited[neighbor] {
					Q = append(Q, neighbor)
				}
			}
		}
		constellations++
	}
	fmt.Printf("a) %d\n", constellations)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(u, v v4) int {
	return abs(u.x-v.x) + abs(u.y-v.y) + abs(u.z-v.z) + abs(u.w-v.w)
}
