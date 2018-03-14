package main

import (
	"fmt"
)

func main() {
	alice := []int{3, 4, 6, 10, 11, 15}
	mine := []int{1, 5, 8, 12, 14, 19}
	fmt.Println(alice)
	fmt.Println(mine)
	merged := mergeArrays(alice, mine)
	fmt.Println(merged)
}

func mergeArrays(a []int, b []int) []int {
	merged := make([]int, len(a)+len(b))
	i := 0
	for len(a) > 0 && len(b) > 0 {
		if a[0] <= b[0] {
			merged[i] = a[0]
			a = a[1:]
		} else {
			merged[i] = b[0]
			b = b[1:]
		}
		i++
	}

	if len(a) > 0 {
		copy(merged[i:], a)
	} else {
		copy(merged[i:], b)
	}
	return merged
}
