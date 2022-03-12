package main

import "fmt"

func fibonacciBottomUp(n int) int {
	memo := map[int]int{}
	// jika belum dihitung
	for i := 0; i <= n; i++ {
		if i <= 0 {
			memo[i] = n
		} else {
			memo[i] = memo[i-1] + memo[i-2]
		}
	}

	return memo[n]

}

func main() {
	fmt.Println(fibonacciBottomUp(0))  //0
	fmt.Println(fibonacciBottomUp(1))  //1
	fmt.Println(fibonacciBottomUp(2))  //1
	fmt.Println(fibonacciBottomUp(3))  //2
	fmt.Println(fibonacciBottomUp(4))  //2
	fmt.Println(fibonacciBottomUp(10)) //2
	fmt.Println(fibonacciBottomUp(40)) //2
	fmt.Println(fibonacciBottomUp(50)) //2

}
