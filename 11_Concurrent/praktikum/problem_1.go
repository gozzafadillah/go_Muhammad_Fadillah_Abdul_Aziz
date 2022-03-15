package main

import (
	"fmt"
	"strings"
)

func Frequency(sentence string) interface{} {
	data := map[string]int{}
	for _, c := range sentence {
		data[string(c)]++
	}
	return data
}

func main() {
	data := "Lorem ipsum dolor, sit amet consectetur adipisicing elit. Obcaecati, deleniti."
	stringSplit := strings.Split(data, " ")
	for i := 0; i <= len(stringSplit)-1; i++ {
		fmt.Println(Frequency(stringSplit[i]))
	}
}
