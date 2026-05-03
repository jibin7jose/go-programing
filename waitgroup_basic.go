package main

import (
	"fmt"
	"sync"
)

/*
-----------------------------------------
File Name: waitgroup_basic.go
Description: Go program demonstrating synchronization using WaitGroup
Author: Jibin Jose
-----------------------------------------

Running Command:
go run waitgroup_basic.go

Output:
Task completed
Main finished
-----------------------------------------
*/

func task(wg *sync.WaitGroup) {
	defer wg.Done() // mark goroutine as done
	fmt.Println("Task completed")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1) // number of goroutines to wait for

	go task(&wg)

	wg.Wait() // wait until all goroutines finish

	fmt.Println("Main finished")

}
