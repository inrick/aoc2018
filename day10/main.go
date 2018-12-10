package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	pos struct{ x, y int }
	v   struct{ x, y int }
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var points []Point
	for sc.Scan() {
		var p Point
		fmt.Sscanf(sc.Text(), "position=<%d, %d> velocity=<%d, %d>",
			&p.pos.x, &p.pos.y, &p.v.x, &p.v.y)
		points = append(points, p)
	}

	fmt.Println("a)")
	t := 0
	for !aligned(points) {
		tick(points)
		t++
	}
	fmt.Printf("b) %d\n", t)
}

func tick(points []Point) {
	for i := range points {
		points[i].pos.x += points[i].v.x
		points[i].pos.y += points[i].v.y
	}
}

func aligned(points []Point) bool {
	minX, minY := 1<<30, 1<<30
	maxX, maxY := 0, 0
	for _, p := range points {
		if minX > p.pos.x {
			minX = p.pos.x
		}
		if maxX < p.pos.x {
			maxX = p.pos.x
		}
		if minY > p.pos.y {
			minY = p.pos.y
		}
		if maxY < p.pos.y {
			maxY = p.pos.y
		}
	}
	lenX := maxX - minX + 1
	lenY := maxY - minY + 1

	if lenY > 10 {
		return false
	}

	grid := make([][]byte, lenY)
	for i := range grid {
		grid[i] = make([]byte, lenX)
	}

	for _, p := range points {
		grid[p.pos.y-minY][p.pos.x-minX] = 3
	}
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] += 32
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
	return true
}
