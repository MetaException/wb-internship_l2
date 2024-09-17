package cmd

import (
	"bufio"
	"dev06/internal/data"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cut [options]",
	Short: "cut - remove sections from each of line",
	Args:  cobra.ExactArgs(0),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().IntP("fields", "f", 1, "выбрать поля (колонки)")
	rootCmd.Flags().StringP("delimiter", "d", "\t", "использовать другой разделитель")
	rootCmd.Flags().BoolP("separated", "s", false, "только строки с разделителем")

	rootCmd.MarkFlagRequired("fields")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		flags, _ := data.NewFlagsParse(cmd)

		fieldNumber := flags.F - 1
		delimiter := flags.D
		isSeparated := flags.S

		sc := bufio.NewScanner(os.Stdin)
		text := make([]string, 0)

		for sc.Scan() {
			text = append(text, sc.Text())
		}

		if err := sc.Err(); err != nil {
			fmt.Println(err)
			return
		}

		writer := bufio.NewWriter(os.Stdout)
		for _, v := range text {
			fields := strings.Split(v, delimiter)

			var toWrite string
			if fieldNumber >= len(fields) {
				toWrite = "\n"

				if isSeparated {
					continue
				}
			} else {
				toWrite = fields[fieldNumber] + "\n"
			}
			writer.WriteString(toWrite)

		}
		writer.Flush()
	}
}
