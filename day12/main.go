package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	initStateStr := sc.Text()[15:]

	init := make([]byte, len(initStateStr))
	for i, c := range []byte(initStateStr) {
		switch c {
		case '#':
			init[i] = 1
		}
	}
	initIxs := make([]int, len(init))
	for i := range initIxs {
		initIxs[i] = i
	}

	sc.Scan() // empty line in between

	// Kind of dumb, but also kind of fun.
	transform := make([]byte, 1<<5)
	for sc.Scan() {
		line := sc.Text()

		var key byte
		for i, c := range []byte(line[0:5]) {
			switch c {
			case '#':
				key |= 1 << (4 - uint(i))
			}
		}

		var pot byte
		switch line[9] {
		case '#':
			pot = 1
		}
		transform[key] = pot
	}

	// a
	state := make([]byte, len(init))
	ixs := make([]int, len(initIxs))
	copy(state, init)
	copy(ixs, initIxs)

	for gen := 0; gen < 20; gen++ {
		N := len(state)
		// If any of the endpoints contain pots we probably need to grow the state
		if state[0]|state[1]|state[2]|state[N-1]|state[N-2]|state[N-3] > 0 {
			state, ixs = grow(state, ixs)
		}
		oldState := make([]byte, len(state))
		copy(oldState, state)
		for i := 2; i < len(oldState)-2; i++ {
			state[i] = transform[pack(oldState[i-2:i+3])]
		}
	}
	sum := 0
	for i := range state {
		if state[i] == 1 {
			sum += ixs[i]
		}
	}
	fmt.Printf("a) %d\n", sum)

	// b
	state = make([]byte, len(init))
	ixs = make([]int, len(initIxs))
	copy(state, init)
	copy(ixs, initIxs)

NextGen:
	for gen := 0; ; gen++ {
		N := len(state)
		if state[0]|state[1]|state[2]|state[N-1]|state[N-2]|state[N-3] > 0 {
			state, ixs = grow(state, ixs)
		}
		oldState := make([]byte, len(state))
		copy(oldState, state)

		for i := 2; i < len(oldState)-2; i++ {
			state[i] = transform[pack(oldState[i-2:i+3])]
		}

		// See if new state is just the old state but shifted
		i, j := 0, 0
		for ; oldState[i] == 0; i++ {
		}
		for ; state[j] == 0; j++ {
		}
		N = len(state)
		for k := i; k < N && k+j-i < N; k++ {
			if oldState[k] != state[k+j-i] {
				continue NextGen
			}
		}
		// Reset indices to what they "should" have been if this was the steady
		// initial state...
		for i := range ixs {
			ixs[i] -= gen + 1
		}
		break
	}
	// ...and move initial steady state 50e9 generations forward.
	sum = 0
	for i := range state {
		if state[i] == 1 {
			sum += ixs[i] + 50000000000
		}
	}

	fmt.Printf("b) %d\n", sum)
}

func pack(b []byte) byte {
	return b[0]<<4 | b[1]<<3 | b[2]<<2 | b[3]<<1 | b[4]
}

// Naively grow state array in both directions
func grow(state []byte, ixs []int) ([]byte, []int) {
	N := len(state)
	M := N * 3
	newState := make([]byte, M)
	newIxs := make([]int, M)

	ixFirst, ixLast := ixs[0], ixs[N-1]
	for i := 0; i < N; i++ {
		newIxs[i] = ixFirst - N + i
		newIxs[N+i] = ixs[i]
		newIxs[2*N+i] = ixLast + i + 1
		newState[N+i] = state[i]
	}

	return newState, newIxs
}
