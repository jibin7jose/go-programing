// -----------------------------------------
// File Name: api.go
// Description: Go program demonstrating simple API creation
// Author: Jibin Jose
// -----------------------------------------

// Running Command:
// go run api.go

// Output:
// Server running at http://localhost:8080/api
//
// Browser / API Response:
// {"message":"API working successfully","status":200}
// -----------------------------------------

package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	response := Response{
		Message: "API working successfully",
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8080", nil)
}
