// -----------------------------------------
// File Name: struct_to_json.go
// Description: Go program demonstrating struct to JSON conversion
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run struct_to_json.go

// Output:
// {"name":"Jibin","age":21,"marks":90}
// -----------------------------------------

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Marks int    `json:"marks"`
}

func main() {

	student := Student{
		Name:  "Jibin",
		Age:   21,
		Marks: 90,
	}

	jsonData, err := json.Marshal(student)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
