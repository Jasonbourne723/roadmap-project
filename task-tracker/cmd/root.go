package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Short: "任务管理",
	Long:  "任务管理器，支持增删改查",
}

func Execute() {
	rootCommand.AddCommand(&listCommand)
	rootCommand.AddCommand(&DeleteCommand)
	rootCommand.AddCommand(&AddCommand)
	rootCommand.AddCommand(&markCommand)
	rootCommand.AddCommand(&updateCommand)

	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
