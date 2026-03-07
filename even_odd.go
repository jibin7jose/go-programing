/*
-----------------------------------------
File Name: evenodd.go
Description: Go program to check whether a number is even or odd
Author: Jibin Jose
-----------------------------------------

Running Command:
go run evenodd.go

Example Output:
Enter a number:
6
6 is an even number.

Another Output:
Enter a number:
7
7 is an odd number.
-----------------------------------------
*/
package main

import "fmt"

func main() {

	// Declare variable
	var number int

	// Ask user for input
	fmt.Println("Enter a number:")
	fmt.Scanln(&number)

	// Check even or odd
	if number%2 == 0 {
		fmt.Println(number, "is an even number.")
	} else {
		fmt.Println(number, "is an odd number.")
	}

}