package main

import (
	"fmt"
	"math"
)

func Frog(frog []int) {
	// variable
	frogNow := 0
	total := 0
	for i := 1; i <= len(frog)-1; i++ {
		// bandingkan
		tujuan1 := int(math.Abs(float64(frog[i] - frog[frogNow])))
		tujuan2 := int(math.Abs(float64(frog[i-1] - frog[frogNow])))
		fmt.Println("tujuan 1 :", tujuan1, "; Tujuan 2 : ", tujuan2)
		if tujuan1 < tujuan2 {
			frogNow = i
			total += tujuan1
		} else {
			frogNow = i - 1
			total += tujuan2
		}

		fmt.Println("total", total)
	}
}

func main() {
	// Frog([]int{10, 30, 40, 20})
	Frog([]int{30, 10, 60, 10, 60, 50})
}

/*

	i				frog									 		frogNow				sum
			jika frog[1] < frog[0] maka frog[i] else frog[0]		frog[0]		total i dengan frogNow
	1				30-10 < 10-10 = frog[0]							0					10-10 = 0
	2				40-10 < 30-10 = frog[1]							1					20
	3				20-30 < 40-30 = frog[3]							3					20+10 = 30
*/
