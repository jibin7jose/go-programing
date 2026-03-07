/*
-----------------------------------------
File Name: sum_n_numbers.go
Description: Go program to calculate the sum of the first N natural numbers
Author: Jibin Jose
-----------------------------------------

Running Command:
go run sum_n_numbers.go

Example Output:
Enter a number:
5
The sum of first 5 numbers is 15
-----------------------------------------
*/

package main

import "fmt"

func main() {

	// Declare variable
	var n int

	// Initialize sum variable
	sum := 0

	// Ask user for input
	fmt.Println("Enter a number:")
	fmt.Scanln(&n)

	// Loop from 1 to n
	for i := 1; i <= n; i++ {

		// Add each number to sum
		sum = sum + i
	}

	// Print result
	fmt.Println("The sum of first", n, "numbers is", sum)
}