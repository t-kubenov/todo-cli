package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"todo-cli/todo"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: todo [add|list|done|delete|deletedone] [arguments]")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: todo add \"Task name\" priority")
			return
		}

		title := args[2]

		var priority todo.Priority
		var err error

		if len(args) == 4 {
			priority, err = todo.ParsePriority(args[3])

			if err != nil {
				fmt.Println(err)
				return
			}
		}

		addTask(title, priority)
	case "list":
		listTasks()
	case "done":
		if len(args) < 3 {
			fmt.Println("Usage: todo done [task ID]")
			return
		}
		id, _ := strconv.Atoi(args[2])
		markTaskDone(id)
	case "delete":
		if len(args) < 3 {
			fmt.Println("Usage: todo delete [task ID]")
			return
		}
		id, _ := strconv.Atoi(args[2])
		deleteTask(id)
	case "deletedone":
		deleteDone()
	default:
		fmt.Println("Unknown command:", command)
	}
}

func addTask(title string, priority todo.Priority) {
	tasks, _ := todo.LoadTasks()
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	task := todo.Task{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		Priority:  priority,
	}

	tasks = append(tasks, task)
	todo.SaveTasks(tasks)

	fmt.Printf("Added task: %s", title)
	if priority.String() != "" {
		fmt.Printf(" (%s)", priority)
	}
	fmt.Println()
}

func listTasks() {
	tasks, _ := todo.LoadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, t := range tasks {
		status := " "
		if t.Completed {
			status = "✔"
		}

		priority := ""
		if t.Priority != todo.None {
			priority = "(" + t.Priority.String() + ")"
		}
		fmt.Printf("[%d] %s %s %s\n", t.ID, status, t.Title, priority)
	}
}

func markTaskDone(id int) {
	tasks, _ := todo.LoadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			todo.SaveTasks(tasks)
			fmt.Println("Marked task as done:", t.Title)
			return
		}
	}
	fmt.Println("Task not found.")
}

func deleteTask(id int) {
	tasks, _ := todo.LoadTasks()
	updated := []todo.Task{}
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		updated = append(updated, t)
	}
	if !found {
		fmt.Println("Task not found.")
		return
	}
	todo.SaveTasks(updated)
	fmt.Println("Deleted task.")
}

func deleteDone() {
	tasks, _ := todo.LoadTasks()
	updated := []todo.Task{}
	found := false

	for _, t := range tasks {
		if !t.Completed {
			updated = append(updated, t)
			continue
		}
		found = true
	}

	todo.SaveTasks(updated)
	if found {
		fmt.Println("Deleted completed tasks.")
	} else {
		fmt.Println("No completed tasks found.")
	}
}
