package cmd

import (
	"fmt"
	"log"
	"roadmap/github-activity/internal/services"

	"github.com/spf13/cobra"
)

var command = cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please input username")
		}
		list, err := services.Srv.Get(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("list: %v\n", list)

	},
}

func Execute() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
