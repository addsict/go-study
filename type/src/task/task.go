package task

import (
	"fmt"
	"user"
)

type Task struct {
	ID     int
	Detail string
	done   bool
	*user.User
}

func NewTask(id int, detail string, assigner *user.User) *Task {
	return &Task{
		ID:     id,
		Detail: detail,
		done:   false,
		User:   assigner,
	}
}

func (task *Task) Finish() {
	task.done = true
}

func (task *Task) String() string {
	return fmt.Sprintf("%d) %s by %s", task.ID, task.Detail, task.FullName())
}

func (task *Task) Print() {
	Print(task)
}

type Stringer interface {
	String() string
}

func Print(stringer Stringer) {
	fmt.Println(stringer.String())
}
