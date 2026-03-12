package main

import "fmt"

/*
-----------------------------------------
File Name: largest_array.go
Description: Go program to find the largest number in an array
Author: Jibin Jose
-----------------------------------------

Running Command:
go run largest_array.go

Output:
The largest number in the array is: 30
-----------------------------------------
*/

func main() {

	numbers := []int{10, 2, 15, 25, 30, 5}
	largest := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}
	}

	fmt.Printf("The largest number in the array is: %d\n", largest)
}