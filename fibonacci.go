package main

import "fmt"

/*
-----------------------------------------
File Name: fibonacci_series.go
Description: Go program to generate Fibonacci series up to n terms
Author: Jibin Jose
-----------------------------------------

Running Command:
go run fibonacci_series.go

Example Output:
Enter number of terms: 5
Fibonacci Series:
0 1 1 2 3
-----------------------------------------
*/

func main() {

	var n int
	first := 0
	second := 1

	fmt.Print("Enter number of terms: ")
	fmt.Scanln(&n)

	fmt.Println("Fibonacci Series:")

	for i := 0; i < n; i++ {
		fmt.Print(first, " ")

		next := first + second
		first = second
		second = next
	}

}
