package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	tasks, err := loadTasks(fileName)
	if err != nil {
		fmt.Println("can't load tasks:", err)
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("write task text, for example: todo add buy milk")
			os.Exit(1)
		}
		text := strings.Join(os.Args[2:], " ")
		tasks = addTask(tasks, text)
		save(tasks)
		fmt.Println("added:", text)

	case "list":
		if len(tasks) == 0 {
			fmt.Println("no tasks yet")
			return
		}
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Text)
		}

	case "done":
		id := getID()
		if !markDone(tasks, id) {
			fmt.Println("task not found:", id)
			os.Exit(1)
		}
		save(tasks)
		fmt.Println("task", id, "is done")

	case "delete":
		id := getID()
		var ok bool
		tasks, ok = deleteTask(tasks, id)
		if !ok {
			fmt.Println("task not found:", id)
			os.Exit(1)
		}
		save(tasks)
		fmt.Println("task", id, "deleted")

	default:
		fmt.Println("unknown command:", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("usage:")
	fmt.Println("  todo add <text>")
	fmt.Println("  todo list")
	fmt.Println("  todo done <id>")
	fmt.Println("  todo delete <id>")
}

// get task id from args
func getID() int {
	if len(os.Args) < 3 {
		fmt.Println("write task id")
		os.Exit(1)
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("id must be a number")
		os.Exit(1)
	}
	return id
}

func save(tasks []Task) {
	err := saveTasks(fileName, tasks)
	if err != nil {
		fmt.Println("can't save tasks:", err)
		os.Exit(1)
	}
}
