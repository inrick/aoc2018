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

	// a
	var ip int
	var regs Regs
	for 0 <= ip && ip < N {
		instr := program[ip]
		regs[ipreg] = ip
		ops[instr[0]](&regs, instr[1], instr[2], instr[3])
		ip = regs[ipreg]
		ip++
	}
	fmt.Printf("a) %d\n", regs[0])

	// b

	// Manually went through the assembly input to reconstruct the code:
	//
	//    int64_t a, b, c, d, e, f;
	//    e = 10551394; // 2*7*167*4513
	//    a = 0;
	//    d = 1;
	//    do {
	//      b = 1;
	//      do {
	//        f = d*b;
	//        if (f == e) { // d*b == e
	//          a += d;
	//        }
	//        ++b;
	//      } while (b <= e);
	//      ++d;
	//    } while (d <= e);
	//    printf("%ld\n", a); // sought value

	// Naive trial division straight through, as n is small enough.
	// Otherwise one could factor n and produce all possible combinations of its
	// factors.
	sum := 0
	n := 10551394
	for d := 1; d <= n; d++ {
		if n%d == 0 {
			sum += d
		}
	}

	fmt.Printf("b) %d\n", sum)
	fmt.Println("(Note: b is not calculated from input)")
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
