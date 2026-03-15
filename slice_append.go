package main

import "fmt"

/*
-----------------------------------------
File Name: slice_append.go
Description: Go program demonstrating how to append elements to a slice
Author: Jibin Jose
-----------------------------------------

Running Command:
go run slice_append.go

Output:
Before Append: [10 20 30]
After Append: [10 20 30 40 50]
-----------------------------------------
*/

func main() {

	numbers := []int{10, 20, 30}

	fmt.Println("Before Append:", numbers)

	numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	fmt.Println("After Append:", numbers)
}
