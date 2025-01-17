package cmd

import (
	"fmt"
	"roadmap/task-tracker/internal/services"

	"github.com/spf13/cobra"
)

var AddCommand = cobra.Command{
	Use:   "add",
	Short: "add new task",
	Long:  "add new task",
	Run: func(cmd *cobra.Command, args []string) {
		des := args[0]
		if err := services.TaskService.Add(des); err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Println("add success")
		}
	},
}
