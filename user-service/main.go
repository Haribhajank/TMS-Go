// user-service/main.go
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Email    string `json:"email"`
}

var users = []User{}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    var newUser User
    json.NewDecoder(r.Body).Decode(&newUser)
    newUser.ID = len(users) + 1
    users = append(users, newUser)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var credentials User
    json.NewDecoder(r.Body).Decode(&credentials)
    for _, user := range users {
        if user.Username == credentials.Username && user.Password == credentials.Password {
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Login successful"))
            return
        }
    }
    w.WriteHeader(http.StatusUnauthorized)
    w.Write([]byte("Invalid credentials"))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func main() {
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/profile", profileHandler)
    http.ListenAndServe(":8080", nil)
}
