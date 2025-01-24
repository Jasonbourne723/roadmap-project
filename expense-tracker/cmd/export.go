package cmd

import (
	"fmt"
	"log"
	"roadmap/expense-tracker/internal/services"

	"github.com/spf13/cobra"
)

var exportCommand = &cobra.Command{
	Use:   "export",
	Long:  "Export",
	Short: "Export",
	Run: func(cmd *cobra.Command, args []string) {
		err := services.RecordSvc.Export()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Export successful")
	},
}
