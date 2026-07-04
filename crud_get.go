// -----------------------------------------
// File Name: crud_get.go
// Description: Go program demonstrating CRUD GET API
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run crud_get.go

// Output:
// Server running at http://localhost:8080/users
//
// Browser / API Response:
// [{"id":1,"name":"Jibin"},{"id":2,"name":"Arun"}]
// -----------------------------------------

package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Jibin"},
	{ID: 2, Name: "Arun"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8080", nil)
}
