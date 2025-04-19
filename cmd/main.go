package main

import (
	"log"
	"net/http"
	"strings"

	"go-task-service/internal/api"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/tasks", api.CreateTask)
	router.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/tasks/")
		if id == "" {
			http.NotFound(w, r)
			return
		}
		api.GetTaskResultByID(w, r, id)
	})

	log.Println("Starting server on :8082")
	if err := http.ListenAndServe(":8082", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
