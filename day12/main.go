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
	copy(state, init)
	startIx := 0

	for gen := 0; gen < 20; gen++ {
		N := len(state)
		// If any of the endpoints contain pots we probably need to grow the state
		if state[0]|state[1]|state[2]|state[N-1]|state[N-2]|state[N-3] > 0 {
			state, startIx = grow(state, startIx)
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
			sum += startIx + i
		}
	}
	fmt.Printf("a) %d\n", sum)

	// b
	state = make([]byte, len(init))
	copy(state, init)
	startIx = 0

NextGen:
	for gen := 0; ; gen++ {
		N := len(state)
		if state[0]|state[1]|state[2]|state[N-1]|state[N-2]|state[N-3] > 0 {
			state, startIx = grow(state, startIx)
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
		// Reset start index to where it would have started if this had been the
		// initial state, then move it 50e9 generations forward.
		startIx = startIx - (gen + 1) + 50e9
		break
	}
	sum = 0
	for i := range state {
		if state[i] == 1 {
			sum += startIx + i
		}
	}

	fmt.Printf("b) %d\n", sum)
}

func pack(b []byte) byte {
	return b[0]<<4 | b[1]<<3 | b[2]<<2 | b[3]<<1 | b[4]
}

// Naively grow state array in both directions
func grow(state []byte, startIx int) ([]byte, int) {
	N := len(state)
	newState := make([]byte, N*3)
	newStartIx := startIx - N

	copy(newState[N:], state)

	return newState, newStartIx
}
