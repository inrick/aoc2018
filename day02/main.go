package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var ids [][]byte
	for sc.Scan() {
		ids = append(ids, []byte(sc.Text()))
	}

	// a
	var wanted [][]byte
	twos, threes := 0, 0
	for _, id := range ids {
		length := len(id)
		sorted := make([]byte, length)
		copy(sorted, id)
		sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

		seen2, seen3 := false, false
		for i := 0; i < length; {
			j := 1
			for ; i+j < length && sorted[i] == sorted[i+j]; j++ {
			}
			switch j {
			case 2:
				seen2 = true
			case 3:
				seen3 = true
			}
			i += j
		}
		if seen2 {
			twos++
		}
		if seen3 {
			threes++
		}
		if seen2 || seen3 {
			wanted = append(wanted, id)
		}
	}
	fmt.Printf("a) %d\n", twos*threes)

	// b
	wlen := len(wanted)
	var result []byte
	for i := 0; i < wlen; i++ {
		for j := i + 1; j < wlen; j++ {
			id1, id2 := wanted[i], wanted[j]
			neq := 0
			for k := range id1 {
				if id1[k] != id2[k] {
					neq++
				}
			}
			if neq == 1 {
				// found it. let's copy the data.
				for k := range id1 {
					if id1[k] == id2[k] {
						result = append(result, id1[k])
					}
				}
				break
			}
		}
	}
	fmt.Printf("b) %s\n", string(result))
}
