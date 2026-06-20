// -----------------------------------------
// File Name: read_file.go
// Description: Go program demonstrating file reading
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run read_file.go

// Output:
// File Content:
// Hello from Go File Handling!
// -----------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	fmt.Println("File Content:")
	fmt.Println(string(data))
}
