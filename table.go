/*
-----------------------------------------
File Name: table.go
Description: Go program to print multiplication table of a given number
Author: Jibin Jose
-----------------------------------------

Running Command:
go run table.go

Example Output:
Enter a number:
5
5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
5 x 4 = 20
5 x 5 = 25
5 x 6 = 30
5 x 7 = 35
5 x 8 = 40
5 x 9 = 45
5 x 10 = 50
-----------------------------------------
*/
package main

import "fmt"

func main() {

	// Declare variable
	var num int

	// Ask user for input
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)

	// Loop from 1 to 10
	for i := 1; i <= 10; i++ {

		// Print multiplication table
		fmt.Printf("%d x %d = %d\n", num, i, num*i)

	}
}