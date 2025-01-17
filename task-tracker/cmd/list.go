package cmd

import (
	"fmt"
	"roadmap/task-tracker/internal/services"

	"github.com/spf13/cobra"
)

func init() {
}

var listCommand = cobra.Command{
	Use:   "list",
	Long:  "get task list",
	Short: "get task list",
	Run: func(cmd *cobra.Command, args []string) {

		var status string
		if len(args) > 0 {
			status = args[0]
		}

		list := services.TaskService.List(status)
		// 打印表头，增加竖线和横线
		fmt.Printf(" %-3s   %-6s  %-26s  %-26s %-30s \n", "ID", "Status", "CreatedAt", "UpdatedAt", "Des")

		// 打印数据行，增加竖线
		for _, d := range list {
			fmt.Printf(" %-3d    %-6s  %-26s  %-26s %-30s\n",
				d.Id, d.Status, d.CreatedAt.Format("2006-01-02"), d.UpdatedAt.Format("2006-01-02"), d.Description)
		}
	},
}
