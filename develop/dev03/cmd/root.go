package cmd

import (
	"dev03/internal/data"
	"dev03/internal/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sort [file]",
	Short: "The sort utility sorts the lines of text",
	Args:  cobra.ExactArgs(1),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("column", "k", 0, "sort by key")
	rootCmd.Flags().BoolP("numeric", "n", false, "sort by numeric")
	rootCmd.Flags().BoolP("reverse", "r", false, "reverse order sorting")
	rootCmd.Flags().BoolP("distinct", "u", false, "unique sort")
	rootCmd.Flags().StringP("output", "o", "", "output path")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		path := args[0]

		flags, err := data.NewFlagsParse(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}

		text, err := utils.ReadText(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		if flags.U {
			text = utils.RemoveDuplicates(text)
		}

		sort.Slice(text, func(i, j int) bool {

			leftWord := text[i]
			rightWord := text[j]

			if flags.K > 0 {
				leftFields := strings.Fields(text[i])
				rightFields := strings.Fields(text[j])

				if len(leftWord) >= flags.K && len(rightFields) >= flags.K {
					leftWord = leftFields[flags.K-1]
					rightWord = rightFields[flags.K-1]
				}
			}

			var comparison bool
			if flags.N {

				leftNumber, err1 := strconv.Atoi(leftWord)
				rightNumber, err2 := strconv.Atoi(rightWord)

				if err1 == nil && err2 == nil {
					comparison = leftNumber < rightNumber
				} else if err1 == nil {
					comparison = false
				} else if err2 == nil {
					comparison = true
				} else {
					comparison = leftWord < rightWord
				}

			} else {
				comparison = leftWord < rightWord
			}

			if flags.R {
				return !comparison
			}
			return comparison
		})

		if flags.OutputPath != "" {
			err = utils.WriteTextToFilePath(text, flags.OutputPath)
		} else {
			err = utils.WriteText(text, os.Stdout)
		}

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
