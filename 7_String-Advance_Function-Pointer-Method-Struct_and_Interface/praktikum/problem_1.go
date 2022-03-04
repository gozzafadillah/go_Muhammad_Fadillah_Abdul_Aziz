package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {

	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		}
	}()

	if len(b) > len(a) {
		a, b = b, a
	}
	cekA := strings.Contains(a, b)
	cekB := strings.Contains(b, a)

	if cekA {
		return b
	} else if cekB {
		return a
	} else {
		panic("Not found!")
	}

}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANG", "KANGOOROO"))
	fmt.Println(Compare("KI", "KIJANNG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
	fmt.Println(Compare("KUKANGAN", "GAN"))
	fmt.Println(Compare("KUKANGAN", "ASW"))
}
