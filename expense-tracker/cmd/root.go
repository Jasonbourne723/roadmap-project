package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Long: "Expense tracker",
}

func Execute() {
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(delCommand)
	rootCmd.AddCommand(listCommand)
	rootCmd.AddCommand(summaryCommand)
	rootCmd.AddCommand(exportCommand)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
