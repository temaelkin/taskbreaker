package main

import (
	"fmt"

	"github.com/temaelkin/taskbreaker/cli"
	"github.com/temaelkin/taskbreaker/storage"
	"github.com/temaelkin/taskbreaker/task"
)

func main() {
	storage := storage.NewJSON("tasks.json")

	tasks, err := storage.Load()
	if err != nil {
		fmt.Println("Load error:", err)
	}

	manager := task.NewManager(tasks)
	cli := cli.NewCli(manager, storage)

	cli.Run()
}
