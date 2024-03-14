package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := taskManager.Tasks
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		fmt.Println("Tasks:")
		for _, t := range tasks {
			completed := "Not Completed"
			if t.Completed {
				completed = "Completed"
			}
			fmt.Printf("- [%d] %s (Due: %s) - %s\n", t.ID, t.Description, t.DueDate.Format("2006-01-02"), completed)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
