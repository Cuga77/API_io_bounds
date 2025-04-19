package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"go-task-service/internal/tasks"

	"github.com/google/uuid"
)

var (
	taskManager = tasks.NewTaskManager()
	taskMutex   sync.Mutex
)

type CreateTaskRequest struct {
}

type CreateTaskResponse struct {
	TaskID string `json:"task_id"`
}

type GetTaskResponse struct {
	TaskID string `json:"task_id"`
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskMutex.Lock()
	defer taskMutex.Unlock()

	taskID := uuid.New().String()
	taskManager.StartTask(taskID) // Таски должны стартовать в отдельной горутине

	response := CreateTaskResponse{TaskID: taskID}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func GetTaskResultByID(w http.ResponseWriter, r *http.Request, id string) {
	result, err := taskManager.GetTaskResult(id)
	if err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	resp := map[string]string{
		"id":     id,
		"result": result,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
