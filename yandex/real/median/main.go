package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, b int
	fmt.Scan(&n, &b)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	count := 0

	for start := 0; start < n; start++ {
		subArray := []int{}

		for end := start; end < n; end++ {
			subArray = append(subArray, a[end])
			if (end-start+1)%2 == 1 {
				sortedSubArray := make([]int, len(subArray))
				copy(sortedSubArray, subArray)
				sort.Ints(sortedSubArray)

				median := sortedSubArray[(end-start)/2]

				if median == b {
					count++

				}
			}
		}
	}

	fmt.Println(count)
}
