package main

import "fmt"

/*
-----------------------------------------
File Name: array_average.go
Description: Go program to calculate the average of numbers in an array
Author: Jibin Jose
-----------------------------------------

Running Command:
go run array_average.go

Output:
The average of the array is: 30.00
-----------------------------------------
*/

func main() {

	numbers := [5]int{10, 20, 30, 40, 50}
	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	average := float64(sum) / float64(len(numbers))

	fmt.Printf("The average of the array is: %.2f\n", average)
}