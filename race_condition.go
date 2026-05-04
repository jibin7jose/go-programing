package main

import (
	"fmt"
	"sync"
)

/*
-----------------------------------------
File Name: race_condition.go
Description: Demonstrates race condition and its fix using Mutex
Author: Jibin Jose
-----------------------------------------

Run:
go run race_condition.go

Check race:
go run -race race_condition.go

Expected Output:
Final Counter: 1000
-----------------------------------------
*/

var counter int = 0
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()   // lock before updating
	counter++
	mu.Unlock() // unlock after updating
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
