package builtin

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func KillCmd(Out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "kill",
		Short: "kill specified proccess",
		Run: func(cmd *cobra.Command, args []string) {
			pid, err := strconv.Atoi(args[0])
			if err != nil {
				return
			}

			if p, err := os.FindProcess(pid); err != nil {
				fmt.Fprintln(Out, err.Error())
			} else {
				if err := p.Kill(); err != nil {
					fmt.Fprintln(Out, err.Error())
				}
			}
		},
	}
}
