// notification-service/main.go
package main

import (
    "encoding/json"
    "net/http"
)

type Notification struct {
    UserID  int    `json:"user_id"`
    Message string `json:"message"`
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
    var notification Notification
    json.NewDecoder(r.Body).Decode(&notification)
    w.Write([]byte("Notification sent"))
}

func main() {
    http.HandleFunc("/notify", notifyHandler)
    http.ListenAndServe(":8082", nil)
}
