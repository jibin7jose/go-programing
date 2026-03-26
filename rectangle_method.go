package main

import "fmt"

/*
-----------------------------------------
File Name: rectangle_method.go
Description: Go program demonstrating struct methods to calculate area of a rectangle
Author: Jibin Jose
-----------------------------------------

Running Command:
go run rectangle_method.go

Output:
Area: 50
-----------------------------------------
*/

type Rectangle struct {
	width  float64
	height float64
}

// Method
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {

	rect := Rectangle{width: 10, height: 5}

	fmt.Println("Area:", rect.area())

}
