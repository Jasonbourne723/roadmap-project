package cmd

import (
	"fmt"
	"log"
	"roadmap/expense-tracker/internal/services"

	"github.com/spf13/cobra"
)

func init() {
	delCommand.PersistentFlags().Int("id", 0, "Id")
}

var delCommand = &cobra.Command{
	Use:   "del",
	Short: "Del",
	Long:  "Del",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			log.Fatal(err)
		}
		if err := services.RecordSvc.Del(int32(id)); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("delete successful")
		}
	},
}
