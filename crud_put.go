// -----------------------------------------
// File Name: crud_put.go
// Description: Go program demonstrating CRUD PUT API
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run crud_put.go

// Test API (cURL):
// curl -X PUT http://localhost:8080/users \
// -H "Content-Type: application/json" \
// -d "{\"id\":2,\"name\":\"Rahul\"}"

// Output:
// Server running at http://localhost:8080/users
//
// API Response:
// [
//   {
//     "id": 1,
//     "name": "Jibin"
//   },
//   {
//     "id": 2,
//     "name": "Rahul"
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

var users = []User{
	{ID: 1, Name: "Jibin"},
	{ID: 2, Name: "Arun"},
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedUser User

	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := range users {
		if users[i].ID == updatedUser.ID {
			users[i].Name = updatedUser.Name
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", updateUser)

	http.ListenAndServe(":8080", nil)
}
