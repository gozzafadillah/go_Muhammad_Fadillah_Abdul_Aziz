package main

import "fmt"

func playDomino(cards [][]int, deck []int) interface{} {
	hasil := [][]int{}

	for i := 0; i <= len(cards)-1; i++ {
		if cards[i][0] == deck[0] || cards[i][0] == deck[1] || cards[i][1] == deck[0] || cards[i][1] == deck[1] {
			hasil = append(hasil, cards[i])
		}

	}
	if len(hasil) >= 2 {
		var total int
		var index int
		var output []int
		var max int
		for i := 0; i < len(hasil); i++ {
			total = hasil[i][0] + hasil[i][1]
			if max <= total {
				total = max
				index = i
			}
			output = append(output, hasil[index]...)
			return output
		}
		return total

	} else {
		return "Tutup Kartu"

	}
}

func main() {
	fmt.Println(playDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}}, []int{3, 6}))
}
