package task

import (
	"encoding/json"
	"os"
	"time"
)

type TaskManager struct {
	Tasks []Task
}

// LoadTasks loads tasks from a JSON file
func (tm *TaskManager) LoadTasks(filename string) error {
	data, err := os.ReadFile(filename) // Use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &tm.Tasks)
	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskManager) SaveTasks(filename string) error {
	tasksJSON, err := json.MarshalIndent(tm.Tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, tasksJSON, 0644) // Use os.WriteFile instead of ioutil.WriteFile
	if err != nil {
		return err
	}

	return nil
}

// AddTask adds a new task to the task manager
func (tm *TaskManager) AddTask(description string, dueDate time.Time) {
	newTask := Task{
		ID:          len(tm.Tasks) + 1,
		Description: description,
		DueDate:     dueDate,
		Completed:   false,
	}
	tm.Tasks = append(tm.Tasks, newTask)
}

// CompleteTask marks a task as completed
func (tm *TaskManager) CompleteTask(taskID int) {
	for i := range tm.Tasks {
		if tm.Tasks[i].ID == taskID {
			tm.Tasks[i].Completed = true
			break
		}
	}
}

// RemoveTask removes a task from the task manager
func (tm *TaskManager) RemoveTask(taskID int) {
	for i, task := range tm.Tasks {
		if task.ID == taskID {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			break
		}
	}
}
