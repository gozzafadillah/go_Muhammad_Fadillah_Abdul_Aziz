package main

import "fmt"

func binnarySearch(arr []int, x int) {
	var hasil int = 0
	var kiri int = 0
	var kanan int = len(arr) - 1

	for kiri <= kanan && hasil == 0 {
		tengah := (kanan + kiri) / 2
		if x < arr[tengah] {
			kanan = tengah - 1
		} else if x > arr[tengah] {
			kiri = tengah + 1
		} else {
			hasil = tengah
		}
	}

	if hasil == 0 {
		fmt.Println("Data : ", x, " Tidak ada")
	} else {
		fmt.Println(hasil)
	}

}

func main() {
	binnarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
	binnarySearch([]int{1, 2, 3, 4, 5, 6}, 2)
}
