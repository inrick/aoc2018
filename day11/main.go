package main

import "fmt"

func main() {
	input := 8772 // "grid serial number"

	const N = 300
	var grid [N][N]int
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			pow := (x+11)*(y+1) + input
			pow *= (x + 11)
			pow %= 1000
			pow /= 100
			pow -= 5
			grid[y][x] = pow
		}
	}

	// a
	var maxSum, maxX, maxY int
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			sum := 0
			for i := 0; i < 3 && y+i < N; i++ {
				for j := 0; j < 3 && x+j < N; j++ {
					sum += grid[y+i][x+j]
				}
			}
			if maxSum < sum {
				maxSum = sum
				maxX = x + 1
				maxY = y + 1
			}
		}
	}
	fmt.Printf("a) %d,%d\n", maxX, maxY)

	// b
	// Fine, be a bit cleverer to avoid having to wait.
	var acc [N + 1][N + 1]int
	for y := 1; y <= N; y++ {
		for x := 1; x <= N; x++ {
			acc[y][x] = grid[y-1][x-1] + acc[y-1][x] + acc[y][x-1] - acc[y-1][x-1]
		}
	}
	var maxSz int
	maxSum, maxX, maxY = 0, 0, 0
	for y := 1; y <= N; y++ {
		for x := 1; x <= N; x++ {
			for sz := 0; x+sz <= N && y+sz <= N; sz++ {
				sum := acc[y-1][x-1] + acc[y+sz][x+sz] - acc[y+sz][x-1] - acc[y-1][x+sz]
				if maxSum < sum {
					maxSum = sum
					maxX = x
					maxY = y
					maxSz = sz + 1
				}
			}
		}
	}
	fmt.Printf("b) %d,%d,%d\n", maxX, maxY, maxSz)
}
