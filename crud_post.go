// -----------------------------------------
// File Name: crud_post.go
// Description: Go program demonstrating CRUD POST API
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run crud_post.go

// Test API Using cURL:
// curl -X POST http://localhost:8080/users \
// -H "Content-Type: application/json" \
// -d "{\"id\":1,\"name\":\"Jibin\"}"

// Output:
// [
//   {
//     "id": 1,
//     "name": "Jibin"
//   }
// ]
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

var users []User

func createUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", createUser)
	http.ListenAndServe(":8080", nil)
}
