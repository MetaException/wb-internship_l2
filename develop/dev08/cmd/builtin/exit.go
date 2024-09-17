package builtin

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func ExitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "exit",
		Short: "exit from shell",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Exited...")
			os.Exit(0)
		},
	}
}
