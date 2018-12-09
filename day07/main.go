package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	from, to byte
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	nodes := make([]int, 26)
	var pairs []Pair
	for sc.Scan() {
		var p Pair
		fmt.Sscanf(
			sc.Text(),
			"Step %c must be finished before step %c can begin.",
			&p.from, &p.to)
		p.from -= 65
		p.to -= 65
		nodes[p.from]++
		nodes[p.to]++
		pairs = append(pairs, p)
	}
	N := 0 // number of nodes
	for _, n := range nodes {
		if n > 0 {
			N++
		}
	}

	edges := adjMatrix(pairs, N)
	steps := make([]byte, 0, N)
	next := nodesWithoutIncomingEdges(edges, N)
	for len(next) > 0 {
		// Reverse sort since we pick from the end and want lowest possible
		// candidate as the next visit.
		sort.Slice(next, func(i, j int) bool { return next[i] > next[j] })
		visit := next[len(next)-1]
		next = next[:len(next)-1]
		steps = append(steps, visit)
		for i := range edges[visit] {
			if edges[visit][i] {
				edges[visit][i] = false
				incoming := false
				for j := 0; j < N; j++ {
					if edges[j][i] {
						incoming = true
						break
					}
				}
				if !incoming {
					next = append(next, byte(i))
				}
			}
		}
	}
	for i := range steps {
		steps[i] += 65
	}
	fmt.Printf("a) %s\n", string(steps))

	// b
	edges = adjMatrix(pairs, N)
	type Step struct {
		t    int
		node byte
	}
	var work []Step
	steps = make([]byte, 0, N)
	next = nodesWithoutIncomingEdges(edges, N)
	t := 0
	for len(next) > 0 || len(work) > 0 {
		sort.Slice(next, func(i, j int) bool { return next[i] > next[j] })
		for len(work) < 5 && len(next) > 0 {
			visit := next[len(next)-1]
			next = next[:len(next)-1]
			work = append(work, Step{t: int(visit) + t + 60, node: visit})
		}
		sort.Slice(work, func(i, j int) bool { return work[i].t < work[j].t })
		for len(work) > 0 && work[0].t <= t {
			w := work[0]
			work = work[1:]
			steps = append(steps, w.node)

			visit := w.node
			for i := range edges[visit] {
				if edges[visit][i] {
					edges[visit][i] = false
					incoming := false
					for j := 0; j < N; j++ {
						if edges[j][i] {
							incoming = true
							break
						}
					}
					if !incoming {
						next = append(next, byte(i))
					}
				}
			}
		}
		t++
	}
	fmt.Printf("b) %d\n", t)
}

func adjMatrix(pairs []Pair, N int) [][]bool {
	edges := make([][]bool, N)
	for i := range edges {
		edges[i] = make([]bool, N)
	}
	for _, p := range pairs {
		edges[p.from][p.to] = true
	}
	return edges
}

func nodesWithoutIncomingEdges(edges [][]bool, N int) []byte {
	var nodes []byte
Loop:
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if edges[j][i] {
				continue Loop
			}
		}
		nodes = append(nodes, byte(i))
	}
	return nodes
}
