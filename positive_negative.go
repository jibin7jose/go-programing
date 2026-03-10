/*
-----------------------------------------
File Name: positive_negative.go
Description: Go program to check whether a number is positive, negative, or zero
Author: Jibin Jose
-----------------------------------------

Running Command:
go run positive_negative.go

Example Output:
Enter a number:
0
0 is zero.

Another Output:
Enter a number:
3
3 is a positive number.

Another Output:
Enter a number:
-3
-3 is a negative number.
-----------------------------------------
*/

package main

import "fmt"

func main() {

	// Declare variable
	var n int

	// Ask user for input
	fmt.Println("Enter a number:")
	fmt.Scan(&n)

	// Check number condition
	if n > 0 {
		fmt.Println(n, "is a positive number.")
	} else if n < 0 {
		fmt.Println(n, "is a negative number.")
	} else {
		fmt.Println(n, "is zero.")
	}
}