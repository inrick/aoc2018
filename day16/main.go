package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	// Note: array is passed by value, no worries about overwriting registers in
	// ops.
	Regs  [4]int
	Instr [4]int // ex. 9 2 1 2 = opcode 9, A=2, B=1, C=2
	Op    = func(r Regs, a, b, c int) Regs

	Test struct {
		before, after Regs
		instr         Instr
	}
)

func main() {
	var ops [16]Op = [...]Op{
		Addr, Addi, Mulr, Muli, Banr, Bani, Borr, Bori,
		Seti, Setr, Gtir, Gtri, Gtrr, Eqir, Eqri, Eqrr,
	}

	var tests []Test
	var program []Instr

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if len(sc.Text()) == 0 {
			continue
		}
		if sc.Text()[0] == 'B' {
			// test case
			var t Test
			fmt.Sscanf(sc.Text(), "Before: [%d, %d, %d, %d]",
				&t.before[0], &t.before[1], &t.before[2], &t.before[3])
			sc.Scan()
			fmt.Sscanf(sc.Text(), "%d %d %d %d",
				&t.instr[0], &t.instr[1], &t.instr[2], &t.instr[3])
			sc.Scan()
			fmt.Sscanf(sc.Text(), "After: [%d, %d, %d, %d]",
				&t.after[0], &t.after[1], &t.after[2], &t.after[3])
			tests = append(tests, t)
		} else {
			// instruction
			var instr Instr
			fmt.Sscanf(sc.Text(), "%d %d %d %d",
				&instr[0], &instr[1], &instr[2], &instr[3])
			program = append(program, instr)
		}
	}

	// a
	threeOrMore := 0
	candidates := make(map[int][]int)
	for _, t := range tests {
		eqops := 0
	OpSearch:
		for i, o := range ops {
			cmp := o(t.before, t.instr[1], t.instr[2], t.instr[3])
			if cmp == t.after {
				eqops++
				lst := candidates[t.instr[0]]
				for _, c := range lst {
					if c == i {
						continue OpSearch
					}
				}
				lst = append(lst, i)
				candidates[t.instr[0]] = lst
			}
		}
		if eqops >= 3 {
			threeOrMore++
		}
	}

	fmt.Printf("a) %d\n", threeOrMore)

	// Very naive method to assign proper op (fine because of small input):
	// Find op with only one candidate. Assign as real then search through and
	// delete it from other candidate lists. Repeat.
	var real [16]Op
AssignReal:
	for k, v := range candidates {
		if len(v) == 1 {
			opnum := v[0]
			real[k] = ops[opnum]
			delete(candidates, k)
			// Purge opnum from all candidate lists
			for k2, v2 := range candidates {
				v3 := v2[:0]
				for _, x := range v2 {
					if x != opnum {
						v3 = append(v3, x)
					}
				}
				candidates[k2] = v3
			}
			goto AssignReal
		}
	}

	// Execute program
	var regs Regs
	for _, instr := range program {
		regs = real[instr[0]](regs, instr[1], instr[2], instr[3])
	}
	fmt.Printf("b) %d\n", regs[0])
}

func Addr(r Regs, a, b, c int) Regs {
	r[c] = r[a] + r[b]
	return r
}

func Addi(r Regs, a, b, c int) Regs {
	r[c] = r[a] + b
	return r
}

func Mulr(r Regs, a, b, c int) Regs {
	r[c] = r[a] * r[b]
	return r
}

func Muli(r Regs, a, b, c int) Regs {
	r[c] = r[a] * b
	return r
}

func Banr(r Regs, a, b, c int) Regs {
	r[c] = r[a] & r[b]
	return r
}

func Bani(r Regs, a, b, c int) Regs {
	r[c] = r[a] & b
	return r
}

func Borr(r Regs, a, b, c int) Regs {
	r[c] = r[a] | r[b]
	return r
}

func Bori(r Regs, a, b, c int) Regs {
	r[c] = r[a] | b
	return r
}

func Setr(r Regs, a, b, c int) Regs {
	r[c] = r[a]
	return r
}

func Seti(r Regs, a, b, c int) Regs {
	r[c] = a
	return r
}

func Gtir(r Regs, a, b, c int) Regs {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func Gtri(r Regs, a, b, c int) Regs {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func Gtrr(r Regs, a, b, c int) Regs {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func Eqir(r Regs, a, b, c int) Regs {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func Eqri(r Regs, a, b, c int) Regs {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func Eqrr(r Regs, a, b, c int) Regs {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}
