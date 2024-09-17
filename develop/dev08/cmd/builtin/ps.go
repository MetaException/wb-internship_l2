package builtin

import (
	"fmt"
	"io"

	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

func PsCmd(Out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "ps",
		Short: "report a snapshot of current proccesses",
		Run: func(cmd *cobra.Command, args []string) {
			if pcs, err := ps.Processes(); err != nil {
				fmt.Fprintln(Out, err)
			} else {
				for _, pc := range pcs {
					fmt.Fprintln(Out, pc.Pid(), pc.Executable())
				}
			}
		},
	}
}
