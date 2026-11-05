package main

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// read tasks from json file
func loadTasks(fileName string) ([]Task, error) {
	var tasks []Task

	data, err := os.ReadFile(fileName)
	if err != nil {
		// no file yet - it's ok, just return empty list
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// write tasks to json file
func saveTasks(fileName string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func addTask(tasks []Task, text string) []Task {
	// new id = biggest id + 1
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	newTask := Task{
		ID:   maxID + 1,
		Text: text,
		Done: false,
	}
	return append(tasks, newTask)
}

func markDone(tasks []Task, id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return true
		}
	}
	return false
}

func deleteTask(tasks []Task, id int) ([]Task, bool) {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, true
		}
	}
	return tasks, false
}
