package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	Regs  [6]int
	Instr [4]int
	Op    = func(r *Regs, a, b, c int)
)

func main() {
	var ops [16]Op = [...]Op{
		Addr, Addi, Mulr, Muli, Banr, Bani, Borr, Bori,
		Seti, Setr, Gtir, Gtri, Gtrr, Eqir, Eqri, Eqrr,
	}

	sc := bufio.NewScanner(os.Stdin)
	var ipreg int
	var program []Instr
	for sc.Scan() {
		if sc.Text()[0] == '#' {
			fmt.Sscanf(sc.Text(), "#ip %d", &ipreg)
		} else {
			var op string
			var instr Instr
			fmt.Sscanf(sc.Text(), "%s %d %d %d", &op, &instr[1], &instr[2], &instr[3])
			switch op {
			case "addr":
				instr[0] = 0
			case "addi":
				instr[0] = 1
			case "mulr":
				instr[0] = 2
			case "muli":
				instr[0] = 3
			case "banr":
				instr[0] = 4
			case "bani":
				instr[0] = 5
			case "borr":
				instr[0] = 6
			case "bori":
				instr[0] = 7
			case "seti":
				instr[0] = 8
			case "setr":
				instr[0] = 9
			case "gtir":
				instr[0] = 10
			case "gtri":
				instr[0] = 11
			case "gtrr":
				instr[0] = 12
			case "eqir":
				instr[0] = 13
			case "eqri":
				instr[0] = 14
			case "eqrr":
				instr[0] = 15
			default:
				panic(sc.Text())
			}
			program = append(program, instr)
		}
	}
	N := len(program)

	// Calculate answer to b) naively by iterating until we reach a value
	// previously seen. That does not guarantee the correct answer as it might
	// occur before entering a periodic phase, but turns out ok in this case.
	seen := make(map[int]byte)
	firstTime := true
	var prev int
	var regs Regs
	for ip := 0; 0 <= ip && ip < N; ip++ {
		// Noted in input: program halts when register 0 equals the value in
		// register 1 when instruction pointer is 28.
		if ip == 28 {
			if firstTime {
				fmt.Printf("a) %d\n", regs[1])
				firstTime = false
			}
			seen[regs[1]] = seen[regs[1]] + 1
			if seen[regs[1]] == 2 {
				break
			}
			prev = regs[1]
		}
		instr := program[ip]
		regs[ipreg] = ip
		ops[instr[0]](&regs, instr[1], instr[2], instr[3])
		ip = regs[ipreg]
	}
	fmt.Printf("b) %d\n", prev)
	fmt.Println("(Note: the code is hardcoded to the input)")
}

func Addr(r *Regs, a, b, c int) {
	r[c] = r[a] + r[b]
}

func Addi(r *Regs, a, b, c int) {
	r[c] = r[a] + b
}

func Mulr(r *Regs, a, b, c int) {
	r[c] = r[a] * r[b]
}

func Muli(r *Regs, a, b, c int) {
	r[c] = r[a] * b
}

func Banr(r *Regs, a, b, c int) {
	r[c] = r[a] & r[b]
}

func Bani(r *Regs, a, b, c int) {
	r[c] = r[a] & b
}

func Borr(r *Regs, a, b, c int) {
	r[c] = r[a] | r[b]
}

func Bori(r *Regs, a, b, c int) {
	r[c] = r[a] | b
}

func Setr(r *Regs, a, b, c int) {
	r[c] = r[a]
}

func Seti(r *Regs, a, b, c int) {
	r[c] = a
}

func Gtir(r *Regs, a, b, c int) {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func Gtri(r *Regs, a, b, c int) {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func Gtrr(r *Regs, a, b, c int) {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func Eqir(r *Regs, a, b, c int) {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func Eqri(r *Regs, a, b, c int) {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func Eqrr(r *Regs, a, b, c int) {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}
