package main

import "fmt"

/*
-----------------------------------------
File Name: palindrome_string.go
Description: Go program to check whether a string is a palindrome
Author: Jibin Jose
-----------------------------------------

Running Command:
go run palindrome_string.go

Example Output:
Enter a string: madam
Palindrome

Enter a string: hello
Not Palindrome
-----------------------------------------
*/

func main() {

	var text string
	reversed := ""

	fmt.Print("Enter a string: ")
	fmt.Scanln(&text)

	for i := len(text) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}

	if text == reversed {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}

}
