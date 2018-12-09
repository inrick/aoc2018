package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var input []int
	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		input = append(input, n)
	}

	_, metadata := visit(input, 0)
	fmt.Printf("a) %d\n", metadata)

	_, value := visit2(input, 0)
	fmt.Printf("b) %d\n", value)
}

func visit(input []int, i int) (j, metadata int) {
	nchildren, nmetadata := input[i], input[i+1]
	i += 2
	for c := 0; c < nchildren; c++ {
		var metadataChild int
		i, metadataChild = visit(input, i)
		metadata += metadataChild
	}
	for m := 0; m < nmetadata; m++ {
		metadata += input[i]
		i++
	}
	j = i
	return
}

func visit2(input []int, i int) (j, value int) {
	nchildren, nmetadata := input[i], input[i+1]
	i += 2
	children := make([]int, 0, nchildren)
	for c := 0; c < nchildren; c++ {
		var valueChild int
		i, valueChild = visit2(input, i)
		children = append(children, valueChild)
	}
	if nchildren == 0 {
		for m := 0; m < nmetadata; m++ {
			value += input[i]
			i++
		}
	} else {
		for m := 0; m < nmetadata; m++ {
			ix := input[i] - 1
			if ix < nchildren {
				value += children[ix]
			}
			i++
		}
	}
	j = i
	return
}
