package main

import (
	"fmt"
	"time"
)

/*
-----------------------------------------
File Name: goroutine_multiple.go
Description: Go program demonstrating multiple goroutines running concurrently
Author: Jibin Jose
-----------------------------------------

Running Command:
go run goroutine_multiple.go

Example Output (order may vary):
Task 1 running
Task 2 running
Task 1 running
Task 2 running
Task 1 running
Task 2 running
Main finished
-----------------------------------------
*/

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, "running")
		time.Sleep(time.Millisecond * 400)
	}
}

func main() {

	go task("Task 1")
	go task("Task 2")

	// wait so goroutines can finish
	time.Sleep(time.Second * 3)

	fmt.Println("Main finished")

}
