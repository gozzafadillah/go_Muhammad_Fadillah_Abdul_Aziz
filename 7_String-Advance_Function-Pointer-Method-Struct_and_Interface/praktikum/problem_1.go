package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var hasil string

	// Cek  panjang dari kedua string lalu cek apakah ada yang sama
	// Jika true
	if len(a) > len(b) && strings.Contains(a, b) {
		for i, _ := range b {
			for j, _ := range a {
				// cek string satu per satu apakah adayang sama
				if strings.Contains(a[j:], b[i:]) {
					hasil = b
				}
			}
		}
	} else if len(a) < len(b) && strings.Contains(b, a) {
		for i, _ := range b {
			for j, _ := range a {
				// cek string satu per satu apakah adayang sama
				if strings.Contains(b[i:], a[j:]) {
					hasil = a
				}
			}
		}
		// jika false
	} else {
		return "tidak ramah match"
	}

	return hasil
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
