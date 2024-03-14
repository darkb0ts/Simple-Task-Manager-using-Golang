package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		description, _ := cmd.Flags().GetString("description")
		dueDateStr, _ := cmd.Flags().GetString("due")

		dueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			fmt.Println("Invalid due date format. Please use YYYY-MM-DD.")
			return
		}

		taskManager.AddTask(description, dueDate)
		fmt.Println("Task added successfully!")
	},
}

func init() {
	addCmd.Flags().StringP("description", "d", "", "Task description (required)")
	addCmd.Flags().StringP("due", "t", time.Now().Format("2006-01-02"), "Due date (format: YYYY-MM-DD)")
	addCmd.MarkFlagRequired("description")
	rootCmd.AddCommand(addCmd)
}
