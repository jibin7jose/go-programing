package main

import "fmt"

/*
-----------------------------------------
File Name: struct_pointer.go
Description: Go program demonstrating use of pointers with structs to modify values
Author: Jibin Jose
-----------------------------------------

Running Command:
go run struct_pointer.go

Output:
Before Update: {Jibin 80}
After Update: {Jibin 95}
-----------------------------------------
*/

type Student struct {
	name  string
	marks int
}

func updateMarks(s *Student) {
	s.marks = 95
}

func main() {

	student := Student{name: "Jibin", marks: 80}

	fmt.Println("Before Update:", student)

	updateMarks(&student)

	fmt.Println("After Update:", student)

}
