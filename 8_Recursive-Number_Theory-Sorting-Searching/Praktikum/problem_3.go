package main

import (
	"fmt"
	"math"
)

func primeNumber(n int) bool {
	// variable
	max_divisor := math.Floor(math.Sqrt(float64(n)))
	toInt := int(n)
	// Jika n == 1 bukan prima
	if n == 1 {
		return false
	}
	for i := 2; i <= int(max_divisor); i++ {
		if toInt%i == 0 {
			return false
		}
	}
	return true
}

func primaSegiEmpat(w, h, s int) int {

	list := []int{}
	n := (w * h)
	var start int = s + 1
	// 2,3,5,7,11,13,15
	for len(list) < n {
		if primeNumber(start) {
			list = append(list, start)
		}
		start++
	}
	fmt.Println(list)
	hasil := 0
	for i := 0; i <= len(list)-1; i++ {
		hasil += list[i]
	}
	return hasil
}

func main() {
	fmt.Println(primaSegiEmpat(2, 3, 13))
	fmt.Println(primaSegiEmpat(5, 2, 1))
}
