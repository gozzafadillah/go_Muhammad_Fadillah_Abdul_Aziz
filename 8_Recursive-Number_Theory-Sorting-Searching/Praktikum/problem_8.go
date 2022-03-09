package main

import "fmt"

type pair struct {
	name  string
	count int
}

func mostAppearTime(items []string) []pair {

	// variable
	data := []pair{}
	item := map[string]int{}

	for _, value := range items {
		item[value]++
	}

	for key, value := range item {
		data = append(data, pair{name: key, count: value})
	}

	return data

}

func main() {
	fmt.Println(mostAppearTime([]string{"Js", "Js", "Golang", "Ruby", "Ruby", "Js", "Js"}))
	fmt.Println(mostAppearTime([]string{"A", "A", "C", "D", "A", "Z", "D", "C", "Z"}))
}
