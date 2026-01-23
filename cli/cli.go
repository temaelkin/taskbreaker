package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/temaelkin/taskbreaker/storage"
	"github.com/temaelkin/taskbreaker/task"
)

type Cli struct {
	manager *task.Manager
	storage storage.TaskStorage
	scanner *bufio.Scanner
}

func NewCli(manager *task.Manager, storage storage.TaskStorage) *Cli {
	return &Cli{
		manager: manager,
		storage: storage,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *Cli) Run() {
	fmt.Println("================")
	fmt.Println()
	fmt.Println("Task<br>eaker")
	fmt.Println("Some motivational motto about getting your things done.")
	fmt.Println("by temaelkin, 2025, v.1.0.0")
	fmt.Println()
	fmt.Println("================")

	fmt.Println()
	fmt.Println("Enter 'help' for some tips.")
	fmt.Println()

	for {
		fmt.Print("> ")
		if !c.scanner.Scan() {
			return
		}

		input := strings.TrimSpace(c.scanner.Text())
		if input == "" {
			continue
		}

		c.handle(input)
	}
}

func (c *Cli) handle(input string) {
	parts := strings.Fields(input)
	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "add":
		c.cmdAdd(args)
	case "done":
		c.cmdDone(args)
	case "all":
		c.cmdAll()
	case "save":
		c.cmdSave()
	case "exit":
		c.cmdExit()
	case "help":
		c.cmdHelp()
	case "shit":
		c.cmdShit()
	default:
		fmt.Println("Unknown command")
	}
}

func (c *Cli) cmdAdd(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: add <task name>")
		return
	}

	name := strings.Join(args, " ")
	c.manager.Add(name)
	fmt.Println("Added:", name)
	fmt.Println("Anything else?")
	fmt.Println()
}

func (c *Cli) cmdDone(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: done <task name>")
		return
	}

	query := strings.Join(args, " ")
	err := c.manager.Done(query)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Marked as done")
}

func (c *Cli) cmdAll() {
	tasks := c.manager.All()
	if len(tasks) == 0 {
		fmt.Println("You have no tasks")
		return
	}

	for _, t := range tasks {
		if t.Done {
			fmt.Println(t.Name, "✅")
		} else {
			fmt.Println(t.Name, "❌")
		}
	}
}

func (c *Cli) cmdSave() {
	err := c.storage.Save(c.manager.All())
	if err != nil {
		fmt.Println("Save error:", err)
		return
	}
	fmt.Println("Saved!")
}

func (c *Cli) cmdExit() {
	_ = c.storage.Save(c.manager.All())
	fmt.Println("Saved!")

	fmt.Println("Bye")
	os.Exit(0)
}

func (c *Cli) cmdHelp() {
	fmt.Println("This is a little task-tracker for a busy fella like you!")
	fmt.Println("It is pretty easy to use.")
	fmt.Println("Here are the commands:")
	fmt.Println("add <task> - if you want to create a new task.")
	fmt.Println("done <task> - if you did something.")
	fmt.Println("all - if you want to see all your tasks.")
	fmt.Println("save - if you need to save your tasks.")
	fmt.Println("exit - if you want to save your tasks & leave.")
	fmt.Println()
}

func (c *Cli) cmdShit() {
	fmt.Println("Finally! A secret command!")
	fmt.Println("I'm sick of being polite!!!")
	fmt.Println("💩")
	fmt.Println("Jerk!")
	fmt.Println()
}
