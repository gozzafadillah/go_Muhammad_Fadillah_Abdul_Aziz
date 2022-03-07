package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {

	key := 1
	//prime := 2 //prime awal [1]
	list := map[int]int{}

	// 1. looping beberapa index prime
	for i := 1; i <= number*5; i++ {
		// 2. dicek apakah data tersebut prime
		if primeNumber(i) {
			// 3. masukan data ke array
			list[key] = i
			key++
		}
	}
	// fmt.Println(list)
	var hasil int
	for key, value := range list {
		if number == key {
			hasil = value
			return hasil
		}
	}

	return hasil

}

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

func main() {
	fmt.Println("hasil : ", primeX(5))
	fmt.Println("hasil : ", primeX(8))
	fmt.Println("hasil : ", primeX(9))
	fmt.Println("hasil : ", primeX(10))
}
