package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/accounts/list", accountsList)
	http.HandleFunc("/users", users)
	http.HandleFunc("/health", health)
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func accountsList(w http.ResponseWriter, r *http.Request) {
	log.Println("New accountsList request")
	accounts := []map[string]interface{}{
		{"name": "account1", "balance": rand.Intn(1000)},
		{"name": "account2", "balance": rand.Intn(1000)},
		{"name": "account3", "balance": rand.Intn(1000)},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(accounts)
	fmt.Fprint(w, string(jsonData))
}

func users(w http.ResponseWriter, r *http.Request) {
	log.Println("New users request")
	users := []map[string]interface{}{
		{"name": "user1", "email": "user1@example.com"},
		{"name": "user2", "email": "user2@example.com"},
		{"name": "user3", "email": "user3@example.com"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(users)
	fmt.Fprint(w, string(jsonData))
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("New health request")
	hostname, _ := os.Hostname()
	health := map[string]interface{}{
		"server": hostname,
		"status": "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(health)
	fmt.Fprint(w, string(jsonData))
}
