package main

import "fmt"

func playWithAsterik(n int) {
	// variable
	var asterik string
	var star string = "*"

	// code
	for i := 1; i <= n; i++ {
		asterik += star
		fmt.Println(asterik)
	}
}

func main() {
	playWithAsterik(5)
}
