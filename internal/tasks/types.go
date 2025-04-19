package tasks

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
	Result string     `json:"result,omitempty"`
}
