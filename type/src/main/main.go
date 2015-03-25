package main

import (
	"fmt"
	"task"
	"user"
)

func main() {
	assigner := user.NewUser("Yuki", "Furuyama")
	task := task.NewTask(1, "buy a notebook", assigner)
	task.Finish()
	fmt.Printf("%+v", task)

	task.Print()
}
