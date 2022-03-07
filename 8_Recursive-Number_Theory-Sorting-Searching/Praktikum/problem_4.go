package main

import (
	"fmt"
	"math"
)

func maxSequent(arr []int) int {
	// variable
	max_so_far := math.MinInt
	max_ending_here := 0

	for i := 0; i < len(arr)-1; i++ {
		max_ending_here = max_ending_here + arr[i]
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
		if max_ending_here < 0 {
			max_ending_here = 0
		}
	}
	return max_so_far
}

func main() {
	fmt.Println(maxSequent([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(maxSequent([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(maxSequent([]int{-7, 21, -11, 3, -1, 11}))
}
