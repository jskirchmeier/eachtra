package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "eachtra",
		Short: "Go on a text based adventure",
		Long:  banner + "\nGo on a text based adventure",
	}

	// flags available to all commands
	debug bool
	quiet bool
)
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Display debugging messages")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Do not display the banner when starting")

}

func initConfig() {

}
