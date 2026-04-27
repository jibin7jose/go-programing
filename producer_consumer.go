package main

import (
	"fmt"
)

/*
-----------------------------------------
File Name: producer_consumer.go
Description: Go program demonstrating producer-consumer pattern using channels
Author: Jibin Jose
-----------------------------------------

Running Command:
go run producer_consumer.go

Output:
Consumed: 1
Consumed: 2
Consumed: 3
Consumed: 4
Consumed: 5
-----------------------------------------
*/

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	for value := range ch {
		fmt.Println("Consumed:", value)
	}
}

func main() {

	ch := make(chan int)

	go producer(ch)
	consumer(ch)

}
