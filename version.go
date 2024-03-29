package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		if !quiet {
			fmt.Println(banner)
		}
		fmt.Println("Version : ", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
