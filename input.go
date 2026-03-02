/*
-----------------------------------------
File Name: input.go
Description: Go program to take user input and print greeting
Author: Jibin Jose
-----------------------------------------

Running Command:
go run input.go

Example Output:
Enter your name:jibin
Hello, jibin!
-----------------------------------------
*/

package main

import "fmt"

func main() {

	// Declare variable
	var name string

	// Ask user for input
	fmt.Print("Enter your name:")

	// Read user input
	fmt.Scanln(&name)

	// Print greeting message
	fmt.Println("Hello, " + name + "!")
}