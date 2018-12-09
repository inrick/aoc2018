package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var coords []Coord
	for sc.Scan() {
		var c Coord
		fmt.Sscanf(sc.Text(), "%d, %d", &c.x, &c.y)
		coords = append(coords, c)
	}

	// Input is small enough that we naively create a large enough grid and
	// calculate all the distances at every point.

	// a
	const N = 500
	var grid [N][N]int
	for i := range grid {
		for j := range grid[i] {
			minDist, minPt := 1<<30, -1
			for k, c := range coords {
				d := dist(c, Coord{i, j})
				if d < minDist {
					minDist = d
					minPt = k
				} else if d == minDist {
					// keep distance if this is not actually closest point
					minPt = -1
				}
			}
			grid[i][j] = minPt
		}
	}

	// anything found on an edge will have inf area
	ignore := make(map[int]bool)
	for i := 0; i < N; i++ {
		ignore[grid[i][0]] = true
		ignore[grid[i][N-1]] = true
		ignore[grid[0][i]] = true
		ignore[grid[N-1][i]] = true
	}

	area := make([]int, len(coords))
	for _, row := range grid {
		for _, elem := range row {
			if elem != -1 {
				area[elem]++
			}
		}
	}
	maxArea := 0
	for i, a := range area {
		if a > maxArea && !ignore[i] {
			maxArea = a
		}
	}
	fmt.Printf("a) %d\n", maxArea)

	// b
	inside := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			d := 0
			for _, c := range coords {
				d += dist(c, Coord{i, j})
			}
			if d < 10000 {
				inside++
			}
		}
	}
	fmt.Printf("b) %d\n", inside)
}

// damn it, go
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Manhattan distance
func dist(p, q Coord) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}
