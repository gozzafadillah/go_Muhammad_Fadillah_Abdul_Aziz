package main

import (
	"fmt"
	"strconv"
)

func findMinMax(arr []int) string {

	var max int = arr[0]
	var min int = arr[0]
	var indexMax int
	var indexMin int
	var hasil string

	for i, value := range arr {
		if value < min {
			min = value
			indexMin = i
		}
		if value > max {
			max = value
			indexMax = i
		}
	}
	hasil = "\nMin : " + strconv.Itoa(min) + " ; Index : " + strconv.Itoa(indexMin) + "\nMax : " + strconv.Itoa(max) + " ; Index : " + strconv.Itoa(indexMax)
	return hasil

}

func main() {
	fmt.Println(findMinMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(findMinMax([]int{2, -5, -4, 22, 7, 7}))
}
