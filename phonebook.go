package main

import "fmt"

/*
------------------------------------------
File Name: phonebook.go
Description: Go program to search a contact in a phonebook using map
Author: Jibin Jose
-----------------------------------------

Running Command:
go run phonebook.go

Example Output:
Enter name to search: Jibin
Phone number: 9999999999

Another Output:
Enter name to search: Rahul
Contact not found
-----------------------------------------
*/

func main() {

	phonebook := make(map[string]string)
	var name string

	// adding contacts
	phonebook["Jibin"] = "9999999999"
	phonebook["Arun"] = "8888888888"

	fmt.Print("Enter name to search: ")
	fmt.Scanln(&name)

	number, found := phonebook[name]

	if found {
		fmt.Println("Phone number:", number)
	} else {
		fmt.Println("Contact not found")
	}

}
