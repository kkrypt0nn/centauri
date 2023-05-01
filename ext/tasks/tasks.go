// Package tasks in an experimental package which serves as an extension that lets users run a function at a specified interval
package tasks

import (
	"sync"
	"time"
)

// Task represents a task that will run a function at a specified interval
type Task struct {
	mutex sync.Mutex
	run   func()
	every time.Duration
	done  chan bool
}

// Start starts the task
func (t *Task) Start() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	ticker := time.NewTicker(t.every)
	go func() {
		t.run()
		for {
			select {
			case <-t.done:
				return
			case <-ticker.C:
				t.run()
			}
		}
	}()
}

// Stop stops the running task
func (t *Task) Stop() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.done <- true
}
