package cmd

import (
	"fmt"
	"roadmap/task-tracker/internal/services"
	"strconv"

	"github.com/spf13/cobra"
)

var markCommand = cobra.Command{
	Use:   "mark",
	Short: "change status",
	Long:  "change status",
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("id error")
			return
		}
		statusList := map[string]bool{"todo": true, "inprogress": true, "done": true}
		if len(args) == 1 {
			services.TaskService.Mark(id, "")
		} else {
			if _, exist := statusList[args[1]]; !exist {
				fmt.Println("status error")
			} else {
				services.TaskService.Mark(id, args[1])
			}
		}
	},
}
