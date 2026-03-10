/*
-----------------------------------------
File Name: largest.go
Description: Go program to find the largest of two numbers
Author: Jibin Jose
-----------------------------------------

Running Command:
go run largest.go

Example Output:
Enter first number: 10
Enter second number: 8
The largest number is: 10
-----------------------------------------
*/

package main

import "fmt"

func main() {

	// Declare variables
	var a, b int

	// Ask user for first number
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	// Ask user for second number
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	// Compare numbers
	if a > b {
		fmt.Println("The largest number is:", a)
	} else {
		fmt.Println("The largest number is:", b)
	}
}