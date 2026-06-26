// -----------------------------------------
// File Name: routes.go
// Description: Go program demonstrating multiple HTTP routes
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run routes.go

// Output:
// Server running at http://localhost:8080
//
// Browser Output:
// http://localhost:8080/         -> Welcome Home
// http://localhost:8080/about    -> About Page
// http://localhost:8080/contact  -> Contact Page
// -----------------------------------------

package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Page")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
