package main

import (
	"fmt"
)

func primeNumber(number int) bool {
	//variable
	index := 0
	var condition bool

	// code
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			index++
		}
	}
	if index == 2 {
		condition = true
	}
	return condition
}

func main() {
	// input
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
