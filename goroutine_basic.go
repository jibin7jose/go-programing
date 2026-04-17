package main

import (
	"fmt"
	"time"
)

/*
-----------------------------------------
File Name: goroutine_basic.go
Description: Go program demonstrating basic goroutine (concurrency)
Author: Jibin Jose
-----------------------------------------

Running Command:
go run goroutine_basic.go

Example Output (order may vary):
Hello from Goroutine
Hello from Main
Hello from Goroutine
Hello from Main
...
-----------------------------------------
*/

func printMessage() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Goroutine")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	go printMessage() // runs in separate goroutine

	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Main")
		time.Sleep(time.Millisecond * 500)
	}

}
