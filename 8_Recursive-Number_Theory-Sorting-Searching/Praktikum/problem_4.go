package main

import "fmt"

func maxSequent(arr []int) int {
	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(arr)-1 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isDone = false
			}
			i++
		}
	}

	var total int
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}

	return total

}

func main() {
	fmt.Println(maxSequent([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSequent([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(maxSequent([]int{-7, 21, -11, 3, -1, 11}))
}
