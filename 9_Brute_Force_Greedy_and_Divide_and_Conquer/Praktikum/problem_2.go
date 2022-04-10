package main

import "fmt"

func moneyCoin(n int) []int {
	coins := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	counter := len(coins) - 1
	kembalian := []int{}

	for n >= 0 && counter >= 0 {
		if n > coins[counter] {
			n = n - coins[counter]
			// fmt.Println(n)
			kembalian = append(kembalian, coins[counter])
		} else {
			counter--
		}

	}
	return kembalian
}

func main() {
	fmt.Println(moneyCoin(123))
	fmt.Println(moneyCoin(543))
	fmt.Println(moneyCoin(7752))
	fmt.Println(moneyCoin(15321))
}

/*
	n		coins		counter
	543		2000		8
	543 	1000		7
	543		500			6 		=> 543 - 500 = 43 hasil
	43		500			6
	43		...			...
	43		20			2		=> 43 - 20 = 23 hasil
	23		20			2		=> 23 - 20	= 3 hasil
	3 		1			0		=> 3 -  1 hasil
	2 		1			0		=> 2 - 1 = 1 hasil
	1		1			0		=> 0 hasil

	[500, 20, 20, 1,1,1]

*/
