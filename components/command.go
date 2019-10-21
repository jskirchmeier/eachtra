package components

import (
	"eachtra/utils"
)

type Command interface {
	Execute(state State, cmd utils.StringQueue) (accepted bool, description string, err error)
}

type CommandFunc func(state State, cmd utils.StringQueue) (accepted bool, description string, err error)

func (f CommandFunc) Execute(state State, cmd utils.StringQueue) (accepted bool, description string, err error) {
	return f(state, cmd)
}

// Commander holds a collection of commands and is itself a command
type Commander struct {
	commands []Command
}

func (cmdr Commander) Execute(state State, cmd utils.StringQueue) (description string, err error) {
	accepted := false
	for _, c := range cmdr.commands {
		accepted, description, err = c.Execute(state, cmd)
		if accepted || err != nil {
			return
		}
	}
	description = "unknown command"
	return
}

func (cmdr *Commander) Add(commands ...Command) {
	for _, c := range commands {
		cmdr.commands = append(cmdr.commands, c)
	}
}

func SimpleCommand(aliases []string, f func(state State, cmd utils.StringQueue) (description string, err error)) Command {
	return CommandFunc(func(state State, cmd utils.StringQueue) (accepted bool, description string, err error) {
		for _, a := range aliases {
			if cmd.Peek() == a {
				accepted = true
				description, err = f(state, cmd)
				return
			}
		}
		return
	})
}

// BaseCommand implements basic functionality to assist in creating commands

// standard commands should be added to all
var (
	// ExitCommand will terminate the current adventure
	ExitCommand = SimpleCommand([]string{"exit", "e", "bye"}, func(state State, cmd utils.StringQueue) (description string, err error) {
		err = ExitAdventure
		return
	})
)
