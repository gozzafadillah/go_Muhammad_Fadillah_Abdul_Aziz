package main

import (
	"fmt"
)

// Cara minimalis
func pow(base, kuadrat int) int {
	var hasil int = 1
	i := 0
	for i < kuadrat {
		if kuadrat%2 == 1 {
			hasil *= base
		}
		kuadrat /= 2
		base *= base
	}

	return hasil

}

// Cara hardcore
func pow2(base, kuadrat int) int {
	// variable
	var y int = 1
	var hasil int = 1

	if kuadrat < 0 {
		base = 1 / base
		kuadrat = -kuadrat
	}

	if kuadrat == 0 {
		y = 1
	}

	for i := 1; i <= kuadrat; i++ {
		if kuadrat%2 == 0 {
			base = base * base
			kuadrat = kuadrat / 2
		} else {
			y = base * y
			base = base * base
			kuadrat = (kuadrat - 1) / 2
		}
	}
	hasil = base * y

	return hasil

}

func main() {
	var base int
	var kuadrat int

	fmt.Print("Masukan angka : ")
	fmt.Scanf("%d \n", &base)
	fmt.Print("Masukan pangkat : ")
	fmt.Scanf("%d \n", &kuadrat)

	fmt.Println("===================>>>>>>>>>>>>>>>> \n")

	fmt.Println("Hasil pangkat dari", base, "^", kuadrat, "=", pow(base, kuadrat))
}
