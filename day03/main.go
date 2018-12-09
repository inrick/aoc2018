package main

import (
	"bufio"
	"fmt"
	"os"
)

type Patch struct {
	id  int
	pos struct{ x, y int }
	dim struct{ x, y int }
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var patches []Patch
	for sc.Scan() {
		var p Patch
		fmt.Sscanf(sc.Text(), "#%d @ %d,%d: %dx%d",
			&p.id, &p.pos.x, &p.pos.y, &p.dim.x, &p.dim.y)
		patches = append(patches, p)
	}

	// fill grid
	var grid [1000][1000]byte
	for _, p := range patches {
		x, y := p.pos.x, p.pos.y
		Nx, Ny := p.dim.x, p.dim.y
		for i := 0; i < Nx; i++ {
			for j := 0; j < Ny; j++ {
				grid[x+i][y+j]++
			}
		}
	}

	// a
	overlap := 0
	for _, row := range grid {
		for _, elem := range row {
			if elem > 1 {
				overlap++
			}
		}
	}
	fmt.Printf("a) %d\n", overlap)

	// b
	var id int
LonerLook:
	for _, p := range patches {
		x, y := p.pos.x, p.pos.y
		Nx, Ny := p.dim.x, p.dim.y
		for i := 0; i < Nx; i++ {
			for j := 0; j < Ny; j++ {
				if grid[x+i][y+j] > 1 {
					continue LonerLook
				}
			}
		}
		id = p.id
		break
	}
	fmt.Printf("b) %d\n", id)
}
