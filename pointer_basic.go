package main

import "fmt"

/*
-----------------------------------------
File Name: pointer_example.go
Description: Go program demonstrating pointer usage (address and dereferencing)
Author: Jibin Jose
-----------------------------------------

Running Command:
go run pointer_example.go

Output:
Value of num: 10
Address of num: 0xc0000140a0   (address may vary)
Pointer value: 0xc0000140a0    (same as address)
Value using pointer: 10
-----------------------------------------
*/

func main() {

	var num int = 10

	var ptr *int = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Pointer value:", ptr)
	fmt.Println("Value using pointer:", *ptr)

}
