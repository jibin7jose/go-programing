/*
-----------------------------------------
File Name: reverse_number.go
Description: Go program to reverse a number entered by the user
Author: Jibin Jose
-----------------------------------------

Running Command:
go run reverse_number.go

Example Output:
Enter a number:
123
Reversed number is: 321
-----------------------------------------
*/

package main

import "fmt"

func main() {

	// Declare variables
	var num, reverse int

	// Ask user for input
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)

	// Loop until number becomes 0
	for num != 0 {

		// Get last digit
		remainder := num % 10

		// Build reversed number
		reverse = reverse*10 + remainder

		// Remove last digit
		num = num / 10
	}

	// Print result
	fmt.Println("Reversed number is:", reverse)
}