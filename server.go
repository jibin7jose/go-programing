// -----------------------------------------
// File Name: server.go
// Description: Go program demonstrating basic HTTP server
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run server.go

// Output:
// Server running at http://localhost:8080
//
// Browser Output:
// Welcome to Go Server!
// -----------------------------------------

package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Server!")
}

func main() {

	http.HandleFunc("/", home)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
