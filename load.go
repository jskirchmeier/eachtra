package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load [name of adventure]",
	Short: "continue an adventure",
	Long:  "continue an adventure",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !quiet {
			fmt.Println(banner)
		}
		return errors.New("not implemented yet, please try again later")
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}


