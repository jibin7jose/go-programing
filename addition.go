/*
-----------------------------------------
File Name: addition.go
Description: Go program to take two numbers from the user and print their sum
Author: Jibin Jose
-----------------------------------------

Running Command:
go run addition.go

Example Output:
Enter the first number:
3
Enter the second number:
4
The sum of 3 and 4 is 7
-----------------------------------------
*/

package main

import "fmt"

func main() {

	// Declare variables
	var a int
	var b int

	// Ask user for first number
	fmt.Println("Enter the first number:")
	fmt.Scanln(&a)

	// Ask user for second number
	fmt.Println("Enter the second number:")
	fmt.Scanln(&b)

	// Calculate sum
	sum := a + b

	// Print result
	fmt.Println("The sum of", a, "and", b, "is", sum)
}
