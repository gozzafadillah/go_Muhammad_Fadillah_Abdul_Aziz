package main

import "fmt"

func simpleEquation(a, b, c int) {
	foundX := 0
	foundY := 0
	foundZ := 0
	errors := ""

	for x := 1; x <= a; x++ {
		for y := 1; y <= b; y++ {
			z := a - (x + y)
			if x+y+z == a && x*y*z == b && (x*x)+(y*y)+(z*z) == c {
				foundX = x
				foundY = y
				foundZ = z
			}
		}
	}
	if foundX == 0 && foundY == 0 && foundZ == 0 {
		errors = "no solution"
		fmt.Println(errors)
	} else {
		fmt.Println(foundX, foundY, foundZ)
	}

}

func main() {
	simpleEquation(6, 6, 14)
	simpleEquation(1, 2, 3)
}
