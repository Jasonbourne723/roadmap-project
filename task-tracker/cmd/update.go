package cmd

import (
	"fmt"
	"roadmap/task-tracker/internal/services"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCommand = cobra.Command{
	Use:   "update",
	Short: "update task description",
	Long:  "update task description",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("args too less")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("id error")
			return
		}

		services.TaskService.Update(id, args[1])
	},
}
