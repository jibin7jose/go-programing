/*
-----------------------------------------
File Name: swap.go
Description: Go program to swap two numbers using multiple assignment
Author: Jibin Jose
-----------------------------------------

Running Command:
go run swap.go

Example Output:
Before swapping: a= 5 b= 10
After swapping: a= 10 b= 5
-----------------------------------------
*/

package main

import "fmt"

func main() {

	// Declare and initialize variables
	var a = 5
	var b = 10

	// Print values before swapping
	fmt.Println("Before swapping: a=", a, "b=", b)

	// Swap values using Go multiple assignment
	a, b = b, a

	// Print values after swapping
	fmt.Println("After swapping: a=", a, "b=", b)
}