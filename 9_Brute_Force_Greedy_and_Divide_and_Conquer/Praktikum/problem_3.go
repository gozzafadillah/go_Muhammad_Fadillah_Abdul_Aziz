package main

import (
	"fmt"
	"sort"
)

func DragonLoowater(Dragon, Knight []int) {
	sort.Ints(Knight)
	sort.Ints(Dragon)

	var hasil int

	for _, value := range Dragon {
		l := len(Knight)
		if k := sort.Search(l,
			func(i int) bool { return Knight[i] >= value }); k < l {
			hasil += Knight[k]
			Knight = Knight[k+1:]
		} else {
			hasil = 0
		}
	}
	if hasil == 0 {
		fmt.Println("Knight Fall!!")
	} else {
		fmt.Println(hasil)

	}
}

func main() {
	DragonLoowater([]int{5, 4}, []int{7, 8, 4})
	DragonLoowater([]int{5, 10}, []int{5})
	DragonLoowater([]int{7, 2}, []int{4, 3, 2, 1})
	DragonLoowater([]int{7, 2}, []int{2, 1, 8, 5})
}
