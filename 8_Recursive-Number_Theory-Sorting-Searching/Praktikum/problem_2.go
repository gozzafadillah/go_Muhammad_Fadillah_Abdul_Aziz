package main

import "fmt"

func fibonacci(n int) int {
	var next int = 0

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		next = (fibonacci(n-1) + fibonacci(n-2))
		return next
	}

}

func main() {
	fmt.Println("0\t >>> \t", fibonacci(0))
	fmt.Println("2\t >>> \t", fibonacci(2))
	fmt.Println("9\t >>> \t", fibonacci(9))
	fmt.Println("10\t >>> \t", fibonacci(10))
	fmt.Println("12\t >>> \t", fibonacci(12))
	fmt.Println("15\t >>> \t", fibonacci(15))
	fmt.Println("20\t >>> \t", fibonacci(20))
	fmt.Println("23\t >>> \t", fibonacci(23))
	fmt.Println("24\t >>> \t", fibonacci(24))
	fmt.Println("30\t >>> \t", fibonacci(30))
}
