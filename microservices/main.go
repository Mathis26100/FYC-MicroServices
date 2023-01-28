package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var servicetype string

func main() {
	switch os.Getenv("SERVICE_TYPE") {
	case "gateway":
		log.Println("Starting as a gateway service")
		servicetype = "gateway"
	case "accounts":
		log.Println("Starting as an accounts service")
		servicetype = "accounts"
	case "users":
		log.Println("Starting as a users service")
		servicetype = "users"
	default:
		log.Println("service unknown")
		os.Exit(1)
	}

	if servicetype == "gateway" {
		accountsUrl := os.Getenv("ACCOUNTS_URL")
		if accountsUrl == "" {
			log.Println("ACCOUNTS_URL not set")
			os.Exit(1)
		}
		usersUrl := os.Getenv("USERS_URL")
		if usersUrl == "" {
			log.Println("USERS_URL not set")
			os.Exit(1)
		}
	}

	http.HandleFunc("/accounts/list", accountsList)
	http.HandleFunc("/users", users)
	http.HandleFunc("/health", health)
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func accountsList(w http.ResponseWriter, r *http.Request) {
	log.Println("New accountsList request")
	if servicetype == "accounts" {
		accounts := []map[string]interface{}{
			{"name": "account1", "balance": rand.Intn(1000)},
			{"name": "account2", "balance": rand.Intn(1000)},
			{"name": "account3", "balance": rand.Intn(1000)},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonData, _ := json.Marshal(accounts)
		fmt.Fprint(w, string(jsonData))
	} else if servicetype == "gateway" {
		url := fmt.Sprintf("http://%s/accounts/list", os.Getenv("ACCOUNTS_URL"))
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonData, _ := json.Marshal(resp.Body)
		fmt.Fprint(w, string(jsonData))
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	log.Println("New users request")
	if servicetype == "users" {
		users := []map[string]interface{}{
			{"name": "user1", "email": "user1@example.com"},
			{"name": "user2", "email": "user2@example.com"},
			{"name": "user3", "email": "user3@example.com"},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonData, _ := json.Marshal(users)
		fmt.Fprint(w, string(jsonData))
	} else if servicetype == "gateway" {
		url := fmt.Sprintf("http://%s/users", os.Getenv("USERS_URL"))
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonData, _ := json.Marshal(resp.Body)
		fmt.Fprint(w, string(jsonData))
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("New health request")
	if servicetype != "gateway" {
		hostname, _ := os.Hostname()
		health := map[string]interface{}{
			"server": hostname,
			"status": "OK",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonData, _ := json.Marshal(health)
		fmt.Fprint(w, string(jsonData))
	} else {
		url := fmt.Sprintf("http://%s/health", os.Getenv("ACCOUNTS_URL"))
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "Accounts service is down")
			return
		}
		url = fmt.Sprintf("http://%s/health", os.Getenv("USERS_URL"))
		resp, _ = http.Get(url)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "Users service is down")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}
