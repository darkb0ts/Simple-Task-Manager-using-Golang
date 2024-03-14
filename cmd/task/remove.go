package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [taskID]",
	Short: "Remove a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID. Please provide a valid integer ID.")
			return
		}

		taskManager.RemoveTask(taskID)
		fmt.Printf("Task %d removed.\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
