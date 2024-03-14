package main

import (
	"fmt"
	"os"

	"task-manager/internal/task"

	"github.com/spf13/cobra"
)

var taskManager = task.TaskManager{}

func init() {
	// Load tasks from file when the program starts
	err := taskManager.LoadTasks("data/tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Simple Task Manager CLI",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Save tasks to file when the program exits
func saveTasksOnExit() {
	err := taskManager.SaveTasks("data/tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func init() {
	// Save tasks to file when the program exits
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Initialize task manager
		taskManager = task.TaskManager{}
		err := taskManager.LoadTasks("data/tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks:", err)
		}
	}

	rootCmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
		saveTasksOnExit()
	}
}
