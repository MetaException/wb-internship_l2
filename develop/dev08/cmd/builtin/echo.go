package builtin

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func EchoCmd(Out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "echo [args]",
		Short: "display a line of text",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(Out, args[0]) //REGEXP quote??
		},
	}
}
