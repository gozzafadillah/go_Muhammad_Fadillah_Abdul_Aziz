package main

import (
	"fmt"
)

func palindrome(input string) bool {
	// variable
	condition := false

	// code
	// dibaca dari index ke 0 dan dibaca terbalik dari index ke n
	for i := 0; i < len(input); i++ {
		if input[i] == input[len(input)-i-1] {
			condition = true
		} else {
			condition = false
		}
	}

	return condition
}

func main() {
	// input
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}
