package main

import (
	"fmt"
	"math"
)

func main() {
	input := 110201
	digits := Digits(input)
	D := len(digits)

	board := []int{3, 7}
	c1, c2 := 0, 1 // Elves

	firstAppears := -1

Outer:
	for len(board) < input+10 || firstAppears == -1 {
		sum := board[c1] + board[c2]
		if sum < 10 {
			board = append(board, sum)
		} else {
			board = append(board, sum/10, sum%10)
		}
		N := len(board)
		c1 = (c1 + 1 + board[c1]) % N
		c2 = (c2 + 1 + board[c2]) % N

		// Every round adds at most two digits. See if there's a match.
		for i := N - D - 2; firstAppears == -1 && 0 <= i && i <= N-D; i++ {
			j := i
			for ; j < N && board[j] != digits[0]; j++ {
			}
			if j > N-D {
				// Not enough digits left for a match
				continue Outer
			}
			k := 0
			for ; k < D && board[j+k] == digits[k]; k++ {
			}
			if k == D {
				firstAppears = j
			}
		}
	}

	fmt.Print("a) ")
	for _, d := range board[input : input+10] {
		fmt.Print(d)
	}
	fmt.Println()
	fmt.Printf("b) %d\n", firstAppears)
}

func Digits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	N := 1 + int(math.Log10(float64(n)))
	digits := make([]int, N)
	for i := range digits {
		digits[N-i-1] = n % 10
		n /= 10
	}
	return digits
}
