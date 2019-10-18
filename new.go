package main

import (
	"bufio"
	"eachtra/adventure"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [name of adventure]",
	Short: "start a new adventure",
	Long:  "start a new adventure",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !quiet {
			fmt.Println(banner)
		}
		if len(args) != 1 {
			return errors.New("name of the adventure must be supplied")
		}
		a, err := adventure.New(args[0])
		if err != nil {
			return err
		}
		fmt.Println(a.Description)
		fmt.Println()
		// blocks until done
		executeCommandLine(a, "Welcome to your adventure ")
		return nil
	},
}

type data struct {
	Turn        int
	Message     string
	Description string
}

const prompt = `{{.Message}}
{{.Description}}

{{printf "%6d" .Turn}} >> `

func executeCommandLine(a *adventure.Adventure, msg string) {
	t, err := template.New("prompt").Parse(prompt)
	if err != nil {
		log.Fatal(err)
	}
	msg = msg + " " + a.Description
	reader := bufio.NewReader(os.Stdin)
	for {
		_ = t.Execute(os.Stdout, data{a.State.Turn, msg, a.State.Location.Introduction()})

		cmd, _ := reader.ReadString('\n')
		msg, err = a.Execute(cmd)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
}
