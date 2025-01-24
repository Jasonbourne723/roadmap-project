package cmd

import (
	"fmt"
	"log"
	"roadmap/expense-tracker/internal/services"

	"github.com/spf13/cobra"
)

func init() {
	addCommand.PersistentFlags().StringP("desc", "d", "", "费用描述")
	addCommand.PersistentFlags().IntP("amount", "a", 0, "金额")
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add",
	Long:  "Add",
	Run: func(cmd *cobra.Command, args []string) {
		desc, err := cmd.Flags().GetString("desc")
		if err != nil {
			log.Fatal(err)
		}
		amount, err := cmd.Flags().GetInt("amount")
		if err != nil {
			log.Fatal(err)
		}
		err = services.RecordSvc.Add(desc, amount)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Add Successful")
		}
	},
}
