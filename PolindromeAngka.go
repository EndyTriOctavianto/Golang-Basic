package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)
	Copy := input
	var Reverse int = 0

	for input > 0 {
		digit := input % 10
		Reverse = Reverse*10 + digit
		input = input / 10
	}
	if Copy == Reverse {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}
