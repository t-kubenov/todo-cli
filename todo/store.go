package todo

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "os"
)

const dbFile = "data/tasks.json"

func LoadTasks() ([]Task, error) {
    if _, err := os.Stat(dbFile); errors.Is(err, os.ErrNotExist) {
        return []Task{}, nil // Return empty if file doesn't exist
    }

    data, err := ioutil.ReadFile(dbFile)
    if err != nil {
        return nil, err
    }

    var tasks []Task
    err = json.Unmarshal(data, &tasks)
    return tasks, err
}

func SaveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(dbFile, data, 0644)
}


