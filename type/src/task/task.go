package task

import (
	"fmt"
)

type Task struct {
	ID     int
	Detail string
	done   bool
}

func NewTask(id int, detail string) *Task {
	return &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
}

func (task *Task) Finish() {
	task.done = true
}

func (task *Task) String() string {
	return fmt.Sprintf("%d) %s", task.ID, task.Detail)
}
