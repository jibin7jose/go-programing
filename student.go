package main

import "fmt"

/*
-----------------------------------------
File Name: struct_student.go
Description: Go program to demonstrate struct usage with user input
Author: Jibin Jose
-----------------------------------------

Running Command:
go run struct_student.go

Example Output:
Enter name: Jibin
Enter age: 22
Enter marks: 85.5

Student Details:
Name: Jibin
Age: 22
Marks: 85.5
-----------------------------------------
*/

type Student struct {
	name  string
	age   int
	marks float64
}

func main() {

	var s Student

	fmt.Print("Enter name: ")
	fmt.Scanln(&s.name)

	fmt.Print("Enter age: ")
	fmt.Scanln(&s.age)

	fmt.Print("Enter marks: ")
	fmt.Scanln(&s.marks)

	fmt.Println("\nStudent Details:")
	fmt.Println("Name:", s.name)
	fmt.Println("Age:", s.age)
	fmt.Println("Marks:", s.marks)

}
