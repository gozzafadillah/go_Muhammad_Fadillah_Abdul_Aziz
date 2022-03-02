package main

import "fmt"

func swap(a, b *int) {
	// penyimpanan sementara
	tempA := *a
	// Swap Logic
	*a = *b
	*b = tempA

}

func main() {
	a := 10
	b := 20

	fmt.Println("number (value) belum diubah  :", a, b)
	fmt.Println("=================///===================")
	swap(&a, &b)

	fmt.Println("number (value) sudah diubah  :", a, b)
	fmt.Println("=================///===================")

}
