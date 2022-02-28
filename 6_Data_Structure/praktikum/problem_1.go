package main

import "fmt"

func tersedia(arr []string, n string) bool {

	for _, i := range arr {
		if i == n {
			return true
		}
	}

	return false
}

func ArrayMerge(ArrayA, ArrayB []string) []string {

	hasil := []string{}

	hasil = append(ArrayA, ArrayB...)

	var newArr []string
	for _, i := range hasil {
		if !tersedia(newArr, i) {
			newArr = append(newArr, i)
		}
	}

	return newArr
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// // ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// // ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"yoshimitsu", "devil jin", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// // ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// // ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))

}
