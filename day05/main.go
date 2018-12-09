package main

import (
	"fmt"
	"os"
)

func main() {
	f := os.Stdin
	st, err := f.Stat()
	if err != nil {
		panic(err)
	}
	N := int(st.Size() - 1) // skip newline in input
	polymer := make([]byte, N)
	read, err := f.Read(polymer)
	if err != nil {
		panic(err)
	}
	if read != N {
		panic("input error")
	}

	// a
	// jmp = backward navigation table, to skip over reacted units
	jmp := make([]int, N+1)
	jmp[0] = -1
	for i := 0; i < N; {
		j, k := i, i+1
		for 0 <= j && k < N && react(polymer[j], polymer[k]) {
			j = jmp[j]
			k++
		}
		jmp[k] = j
		i = k
	}

	count := 0
	for i := jmp[N]; i >= 0; i = jmp[i] {
		count++
	}
	fmt.Printf("a) %d\n", count)

	// b
	minCount := 1 << 30 // large enough
	for skip := byte('A'); skip <= 'Z'; skip++ {
		i := skipOver(skip, polymer, 0)
		jmp[i] = -1
		for i < N {
			j, k := i, skipOver(skip, polymer, i+1)
			for 0 <= j && k < N && react(polymer[j], polymer[k]) {
				j = jmp[j]
				k = skipOver(skip, polymer, k+1)
			}
			jmp[k] = j
			i = k
		}

		count = 0
		for i := jmp[N]; i >= 0; i = jmp[i] {
			count++
		}
		if count < minCount {
			minCount = count
		}
	}
	fmt.Printf("b) %d\n", minCount)
}

func react(x, y byte) bool {
	return (x-32) == y || (x+32) == y
}

func skipOver(skip byte, polymer []byte, i int) int {
	N := len(polymer)
	for i < N && (skip == polymer[i] || skip+32 == polymer[i]) {
		i++
	}
	return i
}
