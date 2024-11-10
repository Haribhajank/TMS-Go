// task-service/main.go
package main

import (
    "encoding/json"
    "net/http"
)

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

var tasks = []Task{}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
    var newTask Task
    json.NewDecoder(r.Body).Decode(&newTask)
    newTask.ID = len(tasks) + 1
    tasks = append(tasks, newTask)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newTask)
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(tasks)
}

func main() {
    http.HandleFunc("/tasks", createTaskHandler)
    http.HandleFunc("/tasks", getTasksHandler)
    http.ListenAndServe(":8081", nil)
}
