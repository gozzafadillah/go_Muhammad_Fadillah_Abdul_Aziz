package main

import (
	"fmt"
)

func main() {
	// input
	studentScore := 63

	// variable
	var grade string

	// code
	if studentScore >= 80 && studentScore <= 100 {
		grade = "A"
	} else if studentScore >= 65 && studentScore <= 79 {
		grade = "B"
	} else if studentScore >= 50 && studentScore <= 64 {
		grade = "C"
	} else if studentScore >= 35 && studentScore <= 49 {
		grade = "D"
	} else if studentScore >= 0 && studentScore <= 34 {
		grade = "E"
	} else {
		grade = "Invalid student score !"
	}

	// output
	fmt.Println("your grade is", grade)
}
