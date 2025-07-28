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
        fmt.Println("Usage: todo [add|list|done|delete] [arguments]")
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
		priority := ""

		if len(args) == 4 {
			priority = args[3]
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
    default:
        fmt.Println("Unknown command:", command)
    }
}

func addTask(title string, priority string) {
    tasks, _ := todo.LoadTasks()
    id := 1
    if len(tasks) > 0 {
        id = tasks[len(tasks)-1].ID + 1
    }

	if priority == "" {
		priority = "medium"
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
    fmt.Println("Added task:", title, "("+priority+")")
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
		if t.Priority != "" {
			priority = "(" + t.Priority + ")"
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
