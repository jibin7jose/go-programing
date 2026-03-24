package main

import "fmt"

/*
-----------------------------------------
File Name: calculator_functions.go
Description: Go program to perform basic arithmetic operations using functions
Author: Jibin Jose
-----------------------------------------

Running Command:
go run calculator_functions.go

Example Output:
Enter first number: 10
Enter second number: 5
Addition: 15
Subtraction: 5
Multiplication: 50
Division: 2
-----------------------------------------
*/

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func main() {

	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	fmt.Println("Addition:", add(a, b))
	fmt.Println("Subtraction:", sub(a, b))
	fmt.Println("Multiplication:", mul(a, b))
	fmt.Println("Division:", div(a, b))

}
