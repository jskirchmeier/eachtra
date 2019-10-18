package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "start the service",
	Long:  "start the service",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !quiet {
			fmt.Println(banner)
		}
		return errors.New("not implemented yet, please try again later")
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}


