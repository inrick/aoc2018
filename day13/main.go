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

func Add(p, q V2) V2 {
	return V2{p.x + q.x, p.y + q.y}
}

func Mult(p, q V2) V2 {
	// (a + bi)(c + di) = ac-bd + (ad+bc)i
	a, b, c, d := p.x, p.y, q.x, q.y
	return V2{a*c - b*d, a*d + b*c}
}

type Cart struct {
	pos, dir  V2
	crossings int // Number of crossings passed
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var track [][]byte
	for sc.Scan() {
		track = append(track, []byte(sc.Text()))
	}

	// Track tokens: | - + / \
	// Cart tokens:  < ^ > v

	// Find carts
	var carts []Cart
	for y := range track {
		for x := range track[y] {
			switch track[y][x] {
			case '<', '^', '>', 'v':
				// Add cart and remove it from track. By design, track is a straight
				// path in carts direction.
				// Note confusion in treating (0,1) as down and (0,-1) as up, to match
				// the track array. Important when rotating further on.
				var dir V2
				switch track[y][x] {
				case '<':
					dir = V2{-1, 0}
					track[y][x] = '-'
				case '^':
					dir = V2{0, -1}
					track[y][x] = '|'
				case '>':
					dir = V2{1, 0}
					track[y][x] = '-'
				case 'v':
					dir = V2{0, 1}
					track[y][x] = '|'
				default:
					panic(nil)
				}
				carts = append(carts, Cart{V2{x, y}, dir, 0})
			}
		}
	}

	// Run carts until all but one has crashed.
	firstCrash := true
	for len(carts) > 1 {
		// Sort carts in track order
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].pos.y < carts[j].pos.y ||
				(carts[i].pos.y == carts[j].pos.y && carts[i].pos.x < carts[j].pos.x)
		})

	NextCart:
		for i := 0; i < len(carts); {
			carts[i].pos = Add(carts[i].pos, carts[i].dir)
			pos := carts[i].pos
			dir := carts[i].dir

			// Note that the rotation is kind of confusing as we consider (0,1) down
			// and (0,-1) up. Be careful!
			var move V2
			switch track[pos.y][pos.x] {
			case '\\':
				if dir.x != 0 {
					move = V2{0, 1}
				} else {
					move = V2{0, -1}
				}
			case '/':
				if dir.x != 0 {
					move = V2{0, -1}
				} else {
					move = V2{0, 1}
				}
			case '+':
				switch carts[i].crossings % 3 {
				case 0:
					move = V2{0, -1}
				case 1:
					move = V2{1, 0}
				case 2:
					move = V2{0, 1}
				}
				carts[i].crossings++
			default:
				// Continue straight ahead
				move = V2{1, 0}
			}
			carts[i].dir = Mult(dir, move)

			// Look for collisions
			for j := range carts {
				if i == j {
					continue
				}
				x0, y0 := carts[i].pos.x, carts[i].pos.y
				x1, y1 := carts[j].pos.x, carts[j].pos.y
				if x0 == x1 && y0 == y1 {
					// Crash!
					if firstCrash {
						fmt.Printf("a) %d,%d\n", x0, y0)
						firstCrash = false
					}
					// Careful, if we remove a cart before us we need to start next
					// iteration one step before current step (as two carts are removed).
					if i < j {
						carts = append(carts[:i], carts[i+1:]...)
						carts = append(carts[:j-1], carts[j:]...)
					} else {
						carts = append(carts[:j], carts[j+1:]...)
						carts = append(carts[:i-1], carts[i:]...)
						i--
					}
					continue NextCart
				}
			}
			i++
		}
	}
	fmt.Printf("b) %d,%d\n", carts[0].pos.x, carts[0].pos.y)
}

// For debugging
func PrintTrack(track [][]byte, carts []Cart) {
	output := make([][]byte, len(track))
	for i, line := range track {
		output[i] = make([]byte, len(line))
		copy(output[i], line)
	}
	for _, c := range carts {
		var b byte
		switch c.dir {
		case V2{-1, 0}:
			b = '<'
		case V2{0, -1}:
			b = '^'
		case V2{1, 0}:
			b = '>'
		case V2{0, 1}:
			b = 'v'
		default:
			panic(nil)
		}
		output[c.pos.y][c.pos.x] = b
	}
	for _, line := range output {
		fmt.Println(string(line))
	}
}
