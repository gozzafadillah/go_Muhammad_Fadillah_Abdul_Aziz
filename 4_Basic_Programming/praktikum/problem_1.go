package main

import (
	"fmt"
)

func main() {
	// input
	var T int = 20
	var r int = 4

	// code
	// const pi
	const pi float32 = 3.14

	// Rumus luas permukaan tabung
	var lp float32 = 2 * pi * float32(r) * (float32(r) + float32(T))

	//output
	fmt.Println(lp)

}
