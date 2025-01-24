package cmd

import (
	"fmt"
	"log"
	"roadmap/expense-tracker/internal/services"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	summaryCommand.PersistentFlags().IntP("year", "y", 0, "年份，默认当年")
	summaryCommand.PersistentFlags().IntP("month", "m", 0, "月份，默认当月")
}

var summaryCommand = &cobra.Command{
	Use:   "summary",
	Short: "summary",
	Long:  "summary",
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
		total, err := services.RecordSvc.Summary(year, month)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("year:%d,month:%d,total:%d\n", year, month, total)

	},
}
