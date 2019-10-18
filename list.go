package main

import (
	"eachtra/adventure"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the available adventures",
	Long:  "list the available adventures",
	Run: func(cmd *cobra.Command, args []string) {

		if !quiet {
			fmt.Println(banner)
		}
		adventures := adventure.List()
		if len(adventures) == 0 {
			fmt.Println("No adventures are available")
		}
		for _, a := range adventures {
			fmt.Println(a)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}


