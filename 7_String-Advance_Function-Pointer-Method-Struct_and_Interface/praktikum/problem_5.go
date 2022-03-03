package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) average() float64 {
	var hasil float64
	// Loop seluruh data score
	for _, value := range s.score {
		// Tambahkan seluruh score
		hasil += float64(value)
	}
	// total score dibagi total index score
	hasil = hasil / float64(len(s.score))

	return hasil
}

func (s Student) min() (min int, name string) {
	// variable
	data := map[string]int{}
	isi := []int{}
	nama := ""
	score := 0
	// memasukan ke dalam looping
	for i := 0; i < len(s.score); i++ {
		data[s.name[i]] = s.score[i]
	}
	// masukan data ke slice isi
	for _, value := range data {
		isi = append(isi, value)
	}

	// sort slice isi
	sort.Ints(isi)
	// mencari data minimal
	for keys, value := range data {
		if value == isi[0] {
			nama = keys
			return value, nama
		}
	}

	return score, nama

}
func (s Student) max() (max int, name string) {
	// variable
	data := map[string]int{}
	isi := []int{}
	nama := ""
	score := 0
	// memasukan ke dalam looping
	for i := 0; i < len(s.score); i++ {
		data[s.name[i]] = s.score[i]
	}
	// masukan data ke slice isi
	for _, value := range data {
		isi = append(isi, value)
	}
	// sort slice isi
	sort.Ints(isi)
	// Mencari data max
	for keys, value := range data {
		if value == isi[len(data)-1] {
			nama = keys
			return value, nama
		}
	}
	return score, nama

}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("input ", i+1, " student name :")
		fmt.Scan(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Print("input ", i+1, " score :")
		fmt.Scan(&score)
		a.score = append(a.score, score)

	}

	fmt.Println("\n\n Average score Students is", a.average())
	scoreMax, nameMax := a.max()
	fmt.Println("Max Score Student is : ", nameMax, "(", scoreMax, ")")
	scoreMin, nameMin := a.min()
	fmt.Println("Min Score Student is : ", nameMin, "(", scoreMin, ")")
}
