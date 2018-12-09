package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var nums []int
	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	// a
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Printf("a) %d\n", sum)

	// b
	sum = 0
	length := len(nums)
	seen := make(map[int]bool)
	for i := 0; !seen[sum]; i++ {
		seen[sum] = true
		sum += nums[i%length]
	}
	fmt.Printf("b) %d\n", sum)
}
