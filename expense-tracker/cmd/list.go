package cmd

import (
	"fmt"
	"log"
	"roadmap/expense-tracker/internal/services"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	listCommand.PersistentFlags().IntP("year", "y", 0, "年份，默认当年")
	listCommand.PersistentFlags().IntP("month", "m", 0, "月份，默认当月")
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List",
	Long:  "List",
	Run: func(cmd *cobra.Command, args []string) {
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			log.Fatal(err)
		}
		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			log.Fatal(err)
		}
		if year == 0 {
			year = time.Now().Year()
		}
		if month == 0 {
			month = int(time.Now().Month())
		}
		list, err := services.RecordSvc.List(year, month)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-3s	%-20s	%-4s	%-8s\n", "id", "description", "amount", "createdAt")
		for _, v := range list {
			fmt.Printf("%-3d	%-20s	%-4d	%-8s\n", v.Id, v.Description, v.Amount, v.CreatedAt.Format("2006-01-02 15:04:05"))
		}
	},
}
