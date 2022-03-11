package main

import (
	"fmt"
	"sort"
)

func maximumBuyProduct(money int, productPrice []int) {

	total := 0
	index := 0

	sort.Ints(productPrice)

	fmt.Println(productPrice)

	for i := 0; i < len(productPrice)-1; i++ {
		total += productPrice[i]
		if total <= money && money > 0 {
			index++
		}
	}
	fmt.Println(index)

}

func main() {

	maximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	maximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	maximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	maximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	maximumBuyProduct(0, []int{10000, 30000})

}
