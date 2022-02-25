package main

import "fmt"

func playWithAsterik(n int) {
	// variable
	var asterik string
	var star string = "* "

	// code
	for i := 1; i <= n; i++ {
		for j := n - i; j >= 1; j-- {
			fmt.Print(" ")
		}
		asterik += star
		fmt.Println(asterik)
	}
}

func main() {
	playWithAsterik(5)
}
