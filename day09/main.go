package main

import (
	"bufio"
	"fmt"
	"os"
)

type Marble struct {
	n           int
	left, right *Marble
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var N, maxPts int
	if sc.Scan() {
		fmt.Sscanf(sc.Text(), "%d players; last marble is worth %d points",
			&N, &maxPts)
	}

	fmt.Printf("a) %d\n", play(N, maxPts))
	fmt.Printf("b) %d\n", play(N, 100*maxPts))
}

func play(N, maxPts int) (highScore int) {
	score := make([]int, N)

	P := 0 // current player
	current := &Marble{0, nil, nil}
	current.left = current
	current.right = current
	for m := 1; m < maxPts; m++ {
		if m%23 == 0 {
			for i := 0; i < 7; i++ {
				current = current.left
			}
			score[P] += m + current.n
			current.left.right = current.right
			current.right.left = current.left
			current = current.right
		} else {
			current = current.right
			newMarble := &Marble{m, current, current.right}
			current.right.left = newMarble
			current.right = newMarble
			current = newMarble
		}
		P = (P + 1) % N
	}
	for _, s := range score {
		if highScore < s {
			highScore = s
		}
	}
	return highScore
}
