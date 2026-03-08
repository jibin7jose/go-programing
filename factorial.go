/*
-----------------------------------------
File Name: factorial.go
Description: Go program to calculate the factorial of a number using a function
Author: Jibin Jose
-----------------------------------------

Running Command:
go run factorial.go

Example Output:
Enter a number:
5
The factorial of 5 is 120
-----------------------------------------
*/

package main

import "fmt"

// Function to calculate factorial
func factorial(n int) int {

	// Initialize result
	result := 1

	// Multiply numbers from 1 to n
	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result
}

func main() {

	// Declare variable
	var n int

	// Ask user for input
	fmt.Println("Enter a number:")
	fmt.Scanln(&n)

	// Call factorial function and print result
	fmt.Println("The factorial of", n, "is", factorial(n))
}