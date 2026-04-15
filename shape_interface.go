package main

import (
	"fmt"
	"math"
)

/*
-----------------------------------------
File Name: shape_interface.go
Description: Go program demonstrating interface (Shape) with Rectangle and Circle implementations
Author: Jibin Jose
-----------------------------------------

Running Command:
go run shape_interface.go

Output:
Area: 50
Area: 28.274333882308138
-----------------------------------------
*/

type Shape interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.area())
}

func main() {

	rect := Rectangle{width: 10, height: 5}
	circle := Circle{radius: 3}

	printArea(rect)
	printArea(circle)

}