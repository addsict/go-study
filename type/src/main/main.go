package main

import (
	"fmt"
	"task"
)

func main() {
	task := task.NewTask(1, "buy a notebook")
	task.Finish()
	fmt.Printf("%+v", task)
}
