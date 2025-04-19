package storage

import (
	"go-task-service/internal/tasks"
	"sync"
)

type MemoryStorage struct {
	mu    sync.RWMutex
	tasks map[string]*tasks.Task
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[string]*tasks.Task),
	}
}

func (s *MemoryStorage) Save(task *tasks.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[task.ID] = task
}

func (s *MemoryStorage) Get(id string) (*tasks.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, exists := s.tasks[id]
	return task, exists
}

func (s *MemoryStorage) GetAll() []*tasks.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	allTasks := make([]*tasks.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}
