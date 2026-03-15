package main

import "fmt"

/*
-----------------------------------------
File Name: count_vowels.go
Description: Go program to count the number of vowels in a string
Author: Jibin Jose
-----------------------------------------

Running Command:
go run count_vowels.go

Example Output:
Enter a string: Hello
Number of vowels: 2
-----------------------------------------
*/

func main() {

	var text string
	count := 0

	fmt.Print("Enter a string: ")
	fmt.Scanln(&text)

	for _, ch := range text {

		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {

			count++
		}

	}

	fmt.Println("Number of vowels:", count)

}
