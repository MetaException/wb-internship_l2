package builtin

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func CdCmd(Out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "cd [path]",
		Short: "cd - switch current directory",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				return
			}

			err := os.Chdir(args[0])
			if err != nil {
				fmt.Fprintln(Out, err.Error())
				return
			}
		},
	}
}
