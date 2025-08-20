package functions

import (
	"encoding/json"
	"fmt"
	"os"
)

const file = "tasks.json"

// function to save to file
func SaveFile(task []Task) error{
	data, err := json.MarshalIndent(task, "", " ")
	if  err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

// function to load from file
func LoadFile()([]Task, error){
	data, err := os.ReadFile(file)
	if err != nil {
		return nil,err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, nil
}

// function to find item by id in file
func FindID(tasks []Task, id int)(*Task, error){
	for i := range tasks{
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, fmt.Errorf("task with id %d not found", id)
}
