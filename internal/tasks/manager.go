package tasks

import (
	"errors"
	"sync"
	"time"
)

type TaskManager struct {
	tasks      map[string]*Task
	mu         sync.Mutex
	taskTicker *time.Ticker
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:      make(map[string]*Task),
		taskTicker: time.NewTicker(1 * time.Second),
	}
}

func (tm *TaskManager) StartTask(id string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task := &Task{
		ID:     id,
		Status: "in progress",
	}
	tm.tasks[id] = task

	go tm.runTask(task)
}

func (tm *TaskManager) runTask(task *Task) {
	time.Sleep(3 * time.Minute)

	tm.mu.Lock()
	defer tm.mu.Unlock()

	task.Status = "completed"
	task.Result = "Task result for ID: " + task.ID
}

func (tm *TaskManager) GetTaskResult(id string) (string, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[id]
	if !exists {
		return "", errors.New("task not found")
	}

	return task.Result, nil
}
