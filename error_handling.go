// -----------------------------------------
// File Name: error_handling.go
// Description: Go program demonstrating error handling
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run error_handling.go

// Output:
// Conversion Error: strconv.Atoi: parsing "123a": invalid syntax
// -----------------------------------------

package main

import (
	"fmt"
	"strconv"
)

func main() {

	input := "123a"

	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Conversion Error:", err)
		return
	}

	fmt.Println("Converted Number:", number)
}
