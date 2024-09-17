package builtin

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func PwdCmd(Out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "pwd",
		Short: "pwd - pring current workdirectory",
		Run: func(cmd *cobra.Command, args []string) {
			currPath, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(Out, err)
				return
			}

			fmt.Fprintln(Out, currPath)
		},
	}
}
