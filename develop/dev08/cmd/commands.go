package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type Command struct {
	*cobra.Command
	CommandMap map[string]*cobra.Command
}

func NewCommand(rootCmd *cobra.Command) *Command {
	return &Command{
		Command:    rootCmd,
		CommandMap: make(map[string]*cobra.Command),
	}
}

func (c *Command) RegisterCommand(command ...*cobra.Command) error {
	for _, cmd := range command {
		if _, ok := c.CommandMap[cmd.Name()]; !ok {
			c.Command.AddCommand(cmd)
			c.CommandMap[cmd.Name()] = cmd
		} else {
			return errors.New("Command is already registered")
		}
	}

	return nil
}
