package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	Range    struct{ from, to int }
	Interval struct{ x, y Range }
	V2       struct{ x, y int }
	Grid     [][]byte
)

var (
	// Useful directions
	Down  = V2{0, 1}
	Left  = V2{-1, 0}
	Right = V2{1, 0}
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var input []Interval
	for sc.Scan() {
		var x, y Range
		switch sc.Text()[0] {
		case 'x':
			fmt.Sscanf(sc.Text(), "x=%d, y=%d..%d", &x.from, &y.from, &y.to)
			x.to = x.from
		case 'y':
			fmt.Sscanf(sc.Text(), "y=%d, x=%d..%d", &y.from, &x.from, &x.to)
			y.to = y.from
		default:
			panic(nil)
		}
		input = append(input, Interval{x, y})
	}

	// Identify bounds
	xmin, xmax := 1<<30, 0
	ymin, ymax := 1<<30, 0
	for _, i := range input {
		if i.x.from < xmin {
			xmin = i.x.from
		}
		if i.x.to > xmax {
			xmax = i.x.to
		}
		if i.y.from < ymin {
			ymin = i.y.from
		}
		if i.y.to > ymax {
			ymax = i.y.to
		}
	}
	xmin, xmax = xmin-1, xmax+1 // Extra room for flow over the edges
	xlen, ylen := xmax-xmin+1, ymax-ymin+1

	// Tokens: . # | ~

	grid := Grid(make([][]byte, ylen))
	for i := range grid {
		grid[i] = make([]byte, xlen)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// Draw basins
	for _, i := range input {
		x, y := i.x, i.y
		if x.from == x.to {
			for j := y.from; j <= y.to; j++ {
				grid[j-ymin][x.from-xmin] = '#'
			}
		} else { // y.from == y.to
			for j := x.from; j <= x.to; j++ {
				grid[y.from-ymin][j-xmin] = '#'
			}
		}
	}

	// Let's do this. Keep flowing until nothing more is added.
	for {
		// Spring of water at x=500, y=0, but remember that grid position is
		// shifted
		st := struct {
			mem      int // State, kind of hacky
			dir, pos V2
		}{
			0, Down, V2{500 - xmin, 0},
		}

		modified := 0
		var forks []V2
		visited := make(map[V2]bool)

	Round:
		for {
			// Use mem = 8 as special state to start revisiting forks in the road.
			if st.mem == 8 {
				if len(forks) == 0 {
					break Round
				}
				st.mem = 0
				st.pos, forks = forks[0], forks[1:]
				st.dir = Right
			}

			// Modify current place in grid
			switch grid.At(st.pos) {
			case '.':
				grid[st.pos.y][st.pos.x] = '|'
				modified++
			case '|':
				if st.mem == 3 {
					grid[st.pos.y][st.pos.x] = '~'
					modified++
				}
			case '~':
				st.mem = 8
				continue Round
			default:
				panic(nil)
			}

			// Are we at the end of the world?
			below := Add(st.pos, Down)
			if below.y == ylen {
				st.mem = 8
				continue Round
			}

			// See if we can fall down
			switch grid.At(below) {
			case '.', '|':
				st.mem = 0
				st.dir = Down
				st.pos = below
				continue Round
			default:
			}

			// Otherwise try to move in current direction
			next := Add(st.pos, st.dir)
			if st.dir == Down {
				switch grid.At(next) {
				case '#', '~':
					st.dir = Left
					if !visited[st.pos] {
						visited[st.pos] = true
						forks = append(forks, st.pos)
					}
				default:
					panic(nil)
				}
			} else {
				switch grid.At(next) {
				case '.', '|':
					st.pos = next
				case '#':
					st.dir.x *= -1
					st.mem++
				case '~':
					st.mem = 8
				default:
					panic(nil)
				}
			}
		}

		if modified == 0 {
			// Finally all flowed out.
			break
		}
	}

	//PrintGrid(grid)

	count := 0
	for _, line := range grid {
		for _, c := range line {
			switch c {
			case '|', '~':
				count++
			}
		}
	}
	fmt.Printf("a) %d\n", count)

	count = 0
	for _, line := range grid {
		for _, c := range line {
			if c == '~' {
				count++
			}
		}
	}
	fmt.Printf("b) %d\n", count)
}

func Add(a, b V2) V2 {
	return V2{a.x + b.x, a.y + b.y}
}

func (grid Grid) At(pos V2) byte {
	return grid[pos.y][pos.x]
}

func PrintGrid(grid Grid) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
