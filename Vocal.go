package main

import "fmt"

func main() {
	fmt.Print("Masukkan Huruf: ")
	var Huruf string
	fmt.Scanln(&Huruf)

	if Huruf == string('a') || Huruf == string('e') || Huruf == string('i') || Huruf == string('e') || Huruf == string('o') || Huruf == string('A') || Huruf == string('I') || Huruf == string('U') || Huruf == string('E') || Huruf == string('O') {
		fmt.Println("VOCAL")
	} else {
		fmt.Println("CONSONANT")
	}
}
