package main

import (
	"fmt"
	"sort"
)

func getMinMax(number ...*int) (min, max int) {
	var data []int
	for _, value := range number {
		data = append(data, *value)
	}

	sort.Ints(data)
	var a, b int = 0, len(data) - 1

	a = data[a]
	b = data[b]

	return a, b
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Print("data terkecil :", min, "\t")
	fmt.Print("data terbesar :", max)
}
