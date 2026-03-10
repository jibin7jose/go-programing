/*
-----------------------------------------
File Name: prime.go
Description: Go program to check whether a number is a prime number
Author: Jibin Jose
-----------------------------------------

Running Command:
go run prime.go

Example Output:
Enter a number:
7
7 is a prime number.

Another Output:
Enter a number:
8
8 is not a prime number.
-----------------------------------------
*/

package main

import "fmt"

func main() {

	// Declare variable
	var n int

	// Flag variable to track prime status
	flag := true

	// Ask user for input
	fmt.Println("Enter a number:")
	fmt.Scan(&n)

	// Check divisibility from 2 to n/2
	for i := 2; i <= n/2; i++ {

		if n%i == 0 {
			flag = false
			break
		}
	}

	// Print result
	if flag {
		fmt.Println(n, "is a prime number.")
	} else {
		fmt.Println(n, "is not a prime number.")
	}
}
