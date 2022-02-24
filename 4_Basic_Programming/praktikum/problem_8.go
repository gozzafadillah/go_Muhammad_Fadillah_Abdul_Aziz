package main

import "fmt"

func cetakTabelPerkalian(number int) {
	// code
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(" ", i*j, "\t")
		}
	}
}

func main() {
	// input
	cetakTabelPerkalian(9)
}
