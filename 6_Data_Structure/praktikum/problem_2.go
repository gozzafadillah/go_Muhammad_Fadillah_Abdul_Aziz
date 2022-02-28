package main

import (
	"fmt"
	"strconv"
	"strings"
)

func angkaMunculSekali(angka string) []int {
	// Variable
	counter := map[string]int{}
	// string di split
	stringSplit := strings.Split(angka, "")
	// menghitung jumlah data yang sama
	for _, item := range stringSplit {
		counter[item]++
	}
	// Membuat variable hasil dengan sesuai ukuran
	hasil := make([]int, 0, len(counter))
	// diubah menjadi int dan masukan ke array
	for key, value := range counter {
		// jika value 1 == true / tidak kembar
		if value == 1 {
			item, err := strconv.Atoi(key)
			if err != nil {
				panic(err)
			}
			hasil = append(hasil, item)
		}
	}

	return hasil
}

func main() {
	var listAngka string
	fmt.Print("List Angka dimasukan masukan (contoh [12345]) : ")
	fmt.Scanf("%s", &listAngka)
	fmt.Println("list angka yang muncul sekali : ", angkaMunculSekali(listAngka))
}
