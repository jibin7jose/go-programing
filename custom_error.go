// -----------------------------------------
// File Name: custom_error.go
// Description: Go program demonstrating custom error handling
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run custom_error.go

// Output:
// Error: insufficient balance
// -----------------------------------------

package main

import (
	"errors"
	"fmt"
)

func withdraw(balance int, amount int) error {

	if amount > balance {
		return errors.New("insufficient balance")
	}

	fmt.Println("Withdrawal successful")
	return nil
}

func main() {

	balance := 1000
	err := withdraw(balance, 1500)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
