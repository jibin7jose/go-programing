package main

import "fmt"

/*
-----------------------------------------
File Name: channel_basic.go
Description: Go program demonstrating basic channel communication between goroutines
Author: Jibin Jose
-----------------------------------------

Running Command:
go run channel_basic.go

Output:
Received: 10
-----------------------------------------
*/

func main() {

	ch := make(chan int)

	go func() {
		ch <- 10 // send value to channel
	}()

	value := <-ch // receive value

	fmt.Println("Received:", value)

}
