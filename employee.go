package main

import "fmt"

/*
-----------------------------------------
File Name: struct_employee.go
Description: Go program to demonstrate struct initialization and display
Author: Jibin Jose
-----------------------------------------

Running Command:
go run struct_employee.go

Output:
Employee Details:
ID: 101
Name: John
Salary: 50000
-----------------------------------------
*/

type Employee struct {
	id     int
	name   string
	salary float64
}

func main() {

	emp := Employee{
		id:     101,
		name:   "John",
		salary: 50000,
	}

	fmt.Println("Employee Details:")
	fmt.Println("ID:", emp.id)
	fmt.Println("Name:", emp.name)
	fmt.Println("Salary:", emp.salary)

}
