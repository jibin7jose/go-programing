package main

import "fmt"

/*
-----------------------------------------
File Name: slice_sum_range.go
Description: Go program to calculate the sum of slice elements using range loop
Author: Jibin Jose
-----------------------------------------

Running Command:
go run slice_sum_range.go

Output:
Numbers: [5 10 15 20]
Sum: 50
-----------------------------------------
*/

func main() {

	numbers := []int{5, 10, 15, 20}

	sum := 0

	for _, value := range numbers {
		sum += value
	}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Sum:", sum)
}
