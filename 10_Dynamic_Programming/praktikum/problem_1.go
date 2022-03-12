package main

import "fmt"

var memo map[int]int = map[int]int{}

func fibonacci(n int) int {
	if nilai, sudahTerhitung := memo[n]; sudahTerhitung {
		return nilai
	}
	// jika belum dihitung
	if n <= 1 {
		memo[n] = n
	} else {
		memo[n] = fibonacci(n-1) + fibonacci(n-2)
	}

	return memo[n]

}

func main() {
	fmt.Println(fibonacci(0))  //0
	fmt.Println(fibonacci(1))  //1
	fmt.Println(fibonacci(2))  //1
	fmt.Println(fibonacci(3))  //2
	fmt.Println(fibonacci(4))  //2
	fmt.Println(fibonacci(10)) //2
	fmt.Println(fibonacci(40)) //2
	fmt.Println(fibonacci(50)) //2

}
