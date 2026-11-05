package main

import (
	"path/filepath"
	"testing"
)

func TestAddTask(t *testing.T) {
	var tasks []Task
	tasks = addTask(tasks, "first")
	tasks = addTask(tasks, "second")

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}
	if tasks[1].ID != 2 {
		t.Errorf("expected id 2, got %d", tasks[1].ID)
	}
}

func TestMarkDone(t *testing.T) {
	tasks := addTask(nil, "test")

	if !markDone(tasks, 1) {
		t.Fatal("task 1 should be found")
	}
	if !tasks[0].Done {
		t.Error("task should be done")
	}
	if markDone(tasks, 99) {
		t.Error("task 99 should not be found")
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := addTask(nil, "a")
	tasks = addTask(tasks, "b")

	tasks, ok := deleteTask(tasks, 1)
	if !ok {
		t.Fatal("task 1 should be deleted")
	}
	if len(tasks) != 1 || tasks[0].Text != "b" {
		t.Errorf("wrong tasks after delete: %v", tasks)
	}
}

func TestSaveAndLoad(t *testing.T) {
	file := filepath.Join(t.TempDir(), "tasks.json")

	tasks := addTask(nil, "hello")
	err := saveTasks(file, tasks)
	if err != nil {
		t.Fatal(err)
	}

	loaded, err := loadTasks(file)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded) != 1 || loaded[0].Text != "hello" {
		t.Errorf("wrong loaded tasks: %v", loaded)
	}
}
