package main

import "fmt"

/*
-----------------------------------------
File Name: reverse_string.go
Description: Go program to reverse a string entered by the user
Author: Jibin Jose
-----------------------------------------

Running Command:
go run reverse_string.go

Example Output:
Enter a string: hello
Reversed string: olleh
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

	fmt.Println("Reversed string:", reversed)

}
