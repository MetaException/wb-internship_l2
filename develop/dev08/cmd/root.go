package cmd

import (
	"bufio"
	"bytes"
	"dev08/cmd/builtin"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rcmd represents the base command when called without any subcommands
var rcmd *Command

func Execute() {
	var out io.Writer
	if len(os.Args) > 1 {
		out = os.Stdout
	} else {
		out = &bytes.Buffer{}
	}
	rcmd = NewCommand(rootCmd(out.(*bytes.Buffer)))
	rcmd.RegisterCommand(builtin.EchoCmd(out), builtin.ExitCmd(), builtin.KillCmd(out), builtin.PsCmd(out), builtin.CdCmd(out), builtin.PwdCmd(out), builtin.HelpCmd(out))

	if err := rcmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd(buf *bytes.Buffer) *cobra.Command {

	return &cobra.Command{
		Use:   "gosh [command]",
		Short: "gosh - go shell",
		Run: func(cmd *cobra.Command, args []string) {

			for {
				currPath, err := os.Getwd()
				if err != nil {
					return
				}
				fmt.Printf("%s> ", currPath)

				sc := bufio.NewReader(os.Stdin)
				line, err := sc.ReadString('\n')
				if err != nil {
					return
				}
				line = trimLine(line)

				inptCmds := strings.Split(line, "|")

				for i, inptCmd := range inptCmds {

					inptCmd = trimLine(inptCmd)
					inpt := strings.Split(inptCmd, " ")

					if i > 0 && i < len(inptCmds) { // pipes
						inpt = append(inpt, buf.String())
						buf.Reset()
					}

					if cmdToRun, ok := rcmd.CommandMap[inpt[0]]; ok {
						cmdToRun.SetArgs(inpt[1:])
						cmdToRun.Run(cmd, inpt[1:])
					} else {
						fmt.Println("Unknown command. Please type help")
					}
				}

				fmt.Print(buf.String())
				buf.Reset()
			}
		},
	}
}

func trimLine(toTrim string) string {
	return strings.Trim(toTrim, " \n\t\r")
}
