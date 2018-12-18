package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var input [][]byte
	for sc.Scan() {
		input = append(input, []byte(sc.Text()))
	}

	// Pad around border so we can always look around
	N := len(input)
	state := make([][]byte, N+2)
	for i := range state {
		state[i] = make([]byte, N+2)
	}
	for i := 1; i <= N; i++ {
		copy(state[i][1:], input[i-1])
	}

	// Tokens:
	// . (open ground)
	// | (tree)
	// # (lumberyard)

	for i := 0; i < 10; i++ {
		Tick(state)
	}

	//PrintYard(state)
	trees, lumber := CountAll(state)
	fmt.Printf("a) %d\n", trees*lumber)

	// Tick on
	seen := make(map[int]int)
	var remaining int
	// Try to find periodicity
	var st struct{ period, count int }
	for i := 10; ; i++ {
		Tick(state)
		trees, lumber := CountAll(state)
		res := trees * lumber
		last := seen[res]
		if st.period == i-last {
			st.count++
		} else {
			st.period = i - last
			st.count = 0
		}
		if st.count == st.period {
			remaining = (1e9 - (i + 1)) % st.period
			break
		}
		seen[res] = i
	}
	for i := 0; i < remaining; i++ {
		Tick(state)
	}
	trees, lumber = CountAll(state)
	fmt.Printf("b) %d\n", trees*lumber)
}

func Tick(state [][]byte) {
	prev := Clone(state)
	for y := range state {
		for x := range state[y] {
			if prev[y][x] == 0 {
				continue
			}
			trees, lumber := CountAdjacent(prev, x, y)
			switch prev[y][x] {
			case '.':
				if trees >= 3 {
					state[y][x] = '|'
				}
			case '|':
				if lumber >= 3 {
					state[y][x] = '#'
				}
			case '#':
				if lumber < 1 || trees < 1 {
					state[y][x] = '.'
				}
			}
		}
	}
}

func CountAll(state [][]byte) (trees, lumber int) {
	for y := range state {
		for x := range state[y] {
			switch state[y][x] {
			case '|':
				trees++
			case '#':
				lumber++
			}
		}
	}
	return
}

func CountAdjacent(state [][]byte, x, y int) (trees, lumber int) {
	xs := []int{x - 1, x, x + 1}
	ys := []int{y - 1, y, y + 1}
	for _, y1 := range ys {
		for _, x1 := range xs {
			if y1 == y && x1 == x {
				continue
			}
			switch state[y1][x1] {
			case '|':
				trees++
			case '#':
				lumber++
			}
		}
	}
	return
}

func Clone(b [][]byte) [][]byte {
	a := make([][]byte, len(b))
	for i := range b {
		a[i] = make([]byte, len(b[i]))
		copy(a[i], b[i])
	}
	return a
}

func PrintYard(state [][]byte) {
	for _, x := range state {
		fmt.Println(string(x))
	}
}
