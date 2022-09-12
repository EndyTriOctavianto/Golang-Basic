package main

import "fmt"

func isPalindrome(str string) bool {
	Reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		Reversed += string(str[i])
	}
	for i := range str {
		if str[i] != Reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Masukkan kata: ")
	var input string
	fmt.Scanln(&input)
	isPalindrome(input)
	if isPalindrome(input) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
