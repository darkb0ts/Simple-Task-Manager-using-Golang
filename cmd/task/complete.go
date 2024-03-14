package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [taskID]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID. Please provide a valid integer ID.")
			return
		}

		taskManager.CompleteTask(taskID)
		fmt.Printf("Task %d marked as completed.\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
