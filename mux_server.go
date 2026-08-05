// -----------------------------------------
// File Name: mux_server.go
// Description: Go program demonstrating Gorilla Mux routing
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go mod init muxdemo
// go get github.com/gorilla/mux
// go run mux_server.go

// Output:
// Server running on :8080
//
// Browser:
// http://localhost:8080/
// Response:
// Welcome to Gorilla Mux!
//
// Browser:
// http://localhost:8080/about
// Response:
// About Page
// -----------------------------------------

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Gorilla Mux!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/about", about).Methods("GET")

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}

