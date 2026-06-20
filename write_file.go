// -----------------------------------------
// File Name: write_file.go
// Description: Go program demonstrating file creation and writing
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run write_file.go

// Output:
// Data written successfully
//
// File Created:
// sample.txt
// Content:
// Hello from Go File Handling!
// -----------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("sample.txt")

	if err != nil {
		fmt.Println("Error creating file")
		return
	}

	defer file.Close()

	file.WriteString("Hello from Go File Handling!")

	fmt.Println("Data written successfully")
}
