package main

import (
	"fmt"
	"math"
)

func primeNumber(n float64) bool {
	// variable
	var condition bool = true
	max_divisor := math.Floor(math.Sqrt(n))
	toInt := int(n)

	if n == 1 {
		condition = false
	}
	for i := 2; i <= int(max_divisor); i++ {
		if toInt%i == 0 {
			condition = false
		} else {
			condition = true
		}
	}
	return condition
}
func main() {
	number := 0.0
	hasil := ""

	fmt.Print("Masukan Bilangan : ")
	fmt.Scanf("%f", &number)

	if primeNumber(number) == true {
		hasil = "Bilangan Prima"
	} else {
		hasil = "Bukan Bilangan Prima"
	}

	fmt.Println("Angka ", int(number), " adalah ", hasil)

}
