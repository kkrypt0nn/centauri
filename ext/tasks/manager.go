package tasks

import (
	"sync"
	"time"
)

// TaskManager represents a manager that holds the different tasks
type TaskManager struct {
	mutex sync.Mutex
	tasks []*Task
}

// NewTaskManager creates a new task manager with empty tasks
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// NewTask creates a new task to run at the specified interval
func (m *TaskManager) NewTask(do func(), every time.Duration) *Task {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	newTask := &Task{
		run:   do,
		every: every,
		done:  make(chan bool),
	}
	m.tasks = append(m.tasks, newTask)
	return newTask
}

// GetTasks returns all the tasks the manager holds
func (m *TaskManager) GetTasks() []*Task {
	return m.tasks
}
