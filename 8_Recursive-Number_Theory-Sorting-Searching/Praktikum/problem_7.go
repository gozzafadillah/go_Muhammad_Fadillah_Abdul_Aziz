package main

import (
	"fmt"
)

func playDomino(cards [][]int, deck []int) {
	hasil := [][]int{}
	sum := 0

	for i := 0; i <= len(cards)-1; i++ {
		counter := 0
		for j := 0; j <= len(deck)-1; j++ {
			if cards[i][counter] == deck[j] {
				hasil = append(hasil, cards[i:][counter:]...)
				sum += cards[i][j]
			}
		}
		counter++
		if counter > len(deck)-1 {
			counter = 0
		}
	}

	fmt.Println(sum)
	fmt.Println(hasil)

}

func main() {
	playDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{3, 3}}, []int{4, 3})
	playDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}}, []int{3, 6})
}
