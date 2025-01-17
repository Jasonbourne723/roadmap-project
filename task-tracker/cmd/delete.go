package cmd

import (
	"fmt"
	"roadmap/task-tracker/internal/services"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
}

var DeleteCommand = cobra.Command{
	Use:   "del",
	Long:  "delete a task",
	Short: "delete a task ",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(fmt.Errorf("id error:%w", err))
		}
		services.TaskService.Delete(id)
		fmt.Println("done!")
	},
}
