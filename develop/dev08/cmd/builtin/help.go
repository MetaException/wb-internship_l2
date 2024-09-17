package builtin

import (
	"io"

	"github.com/spf13/cobra"
)

func HelpCmd(Out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "prints help message",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Root().Help()
		},
	}
}
