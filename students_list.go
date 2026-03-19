package main

import "fmt"

/*
-----------------------------------------
File Name: students_list.go
Description: Go program to demonstrate slice of structs and display student details
Author: Jibin Jose
-----------------------------------------

Running Command:
go run students_list.go

Output:
Student List:

Name: Jibin
Age: 21
Marks: 85.5
-------------------
Name: Arun
Age: 22
Marks: 78
-------------------
Name: Meera
Age: 20
Marks: 92.3
-------------------
-----------------------------------------
*/

type Student struct {
	name  string
	age   int
	marks float64
}

func main() {

	students := []Student{
		{"Jibin", 21, 85.5},
		{"Arun", 22, 78.0},
		{"Meera", 20, 92.3},
	}

	fmt.Println("Student List:\n")

	for _, s := range students {
		fmt.Println("Name:", s.name)
		fmt.Println("Age:", s.age)
		fmt.Println("Marks:", s.marks)
		fmt.Println("-------------------")
	}

}
