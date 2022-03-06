package main

import (
	"fmt"
	"strings"
)

type Student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (d *Student) Encode() string {
	var nameEncode = ""

	for _, value := range d.name {
		if strings.ToUpper(string(value)) == string(value) {
			nameEncode += string(((value + rune(9) - 65) % 26) + 65)
		} else {
			nameEncode += string(((value + rune(9) - 97) % 26) + 97)
		}
	}

	return nameEncode
}

func (d *Student) Decode() string {
	var namaDecode = ""

	for _, value := range d.name {
		if strings.ToUpper(string(value)) == string(value) {
			namaDecode += string(((value + rune(17) - 65) % 26) + 65)
		} else {
			namaDecode += string(((value + rune(17) - 97) % 26) + 97)
		}
	}

	return namaDecode
}

func main() {
	var pilihan int
	var student Student
	var c = &student

	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		}
	}()

	// buat display
	fmt.Println("Pilih Menu")
	fmt.Println("[1.] Encrypt")
	fmt.Println("[2.] Decrypt")
	fmt.Print("Pilihan anda : ")
	fmt.Scanf("%d", &pilihan)

	if pilihan == 1 {
		// input
		fmt.Print("\n Input Student`s name : ")
		fmt.Scan(&student.name)
		fmt.Print("\n Encode of Student`s name "+student.name+" is : ", c.Encode()+"\n")
		fmt.Println("============================ Batas ==============================")
		student.Encode()
	} else if pilihan == 2 {
		// fungsi decrypt
		fmt.Print("\n Input Student`s encode name : ")
		fmt.Scan(&student.name)
		fmt.Print("\n Decode of Student`s name "+student.name+" is : ", c.Decode()+"\n")
		fmt.Println("============================ Batas ==============================")
	} else {
		panic("Not found!!")
	}
}
