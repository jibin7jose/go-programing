package main

import (
	"fmt"
	"sync"
)

/*
-----------------------------------------
File Name: waitgroup_multiple.go
Description: Go program demonstrating multiple goroutines synchronized using WaitGroup
Author: Jibin Jose
-----------------------------------------

Running Command:
go run waitgroup_multiple.go

Example Output (order may vary):
Worker 1 done
Worker 3 done
Worker 2 done
All workers finished
-----------------------------------------
*/

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker", id, "done")
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All workers finished")

}
