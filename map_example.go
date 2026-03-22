package main

import "fmt"

/*
-----------------------------------------
File Name: map_example.go
Description: Go program to demonstrate map creation, insertion, and access
Author: Jibin Jose
-----------------------------------------

Running Command:
go run map_example.go

Output:
Student Marks: map[Arun:78 Jibin:85 Meera:92]
Marks of Arun: 78
-----------------------------------------
*/

func main() {

	// creating a map
	studentMarks := make(map[string]int)

	// adding values
	studentMarks["Jibin"] = 85
	studentMarks["Arun"] = 78
	studentMarks["Meera"] = 92

	fmt.Println("Student Marks:", studentMarks)

	// accessing value
	fmt.Println("Marks of Arun:", studentMarks["Arun"])

}