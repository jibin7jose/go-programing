// -----------------------------------------
// File Name: crud_delete.go
// Description: Go program demonstrating CRUD DELETE API
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run crud_delete.go

// Test API (cURL):
// curl -X DELETE http://localhost:8080/users \
// -H "Content-Type: application/json" \
// -d "{\"id\":2}"

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
//     "id": 3,
//     "name": "Meera"
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
	{ID: 3, Name: "Meera"},
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE allowed", http.StatusMethodNotAllowed)
		return
	}

	var request User

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == request.ID {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", deleteUser)

	http.ListenAndServe(":8080", nil)
}
