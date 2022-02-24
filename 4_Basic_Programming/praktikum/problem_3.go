package main

import (
	"fmt"
)

func main() {
	// Deklarasi Variable
	var bilangan int

	// Input by scanf
	fmt.Print("Masukan Bilangan : ")
	fmt.Scanf("%d", &bilangan)

	// code
	for i := 1; i <= bilangan; i++ {
		if bilangan%i != 0 {
			continue
		}
		fmt.Println(i)

	}

}
