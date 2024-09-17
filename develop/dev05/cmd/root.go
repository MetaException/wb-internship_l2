package cmd

import (
	"bufio"
	"dev05/internal/data"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep [options] [patterns]",
	Short: "grep searches for PATTERNS in each FILE.",
	Args:  cobra.ExactArgs(1),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("ignore-case", "i", false, "игнорировать регистр")
	rootCmd.Flags().BoolP("invert", "v", false, "инвертированный поиск")
	rootCmd.Flags().BoolP("fixed", "F", false, "точное совпадение со строкой, не паттерн")
	rootCmd.Flags().IntP("after", "A", -1, "печать +N строк после совпадения")
	rootCmd.Flags().IntP("before", "B", -1, "печать +N строк до совпадения")
	rootCmd.Flags().IntP("context", "C", -1, "печать +-строк вокруг совпадения")
	rootCmd.Flags().BoolP("count", "c", false, "вывести количество строк")
	rootCmd.Flags().BoolP("line-num", "n", false, "напечать номер строки")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		pattern := args[0]
		//path := args[1]

		flags, _ := data.NewFlagsParse(cmd)

		sc := bufio.NewScanner(os.Stdin)
		text := make([]string, 0)

		for sc.Scan() {
			text = append(text, sc.Text())
		}

		if err := sc.Err(); err != nil {
			return
		}

		count := 0
		writer := bufio.NewWriter(os.Stdout)

		lastAcceptedIndex := -1
		needWriteCount := 0

		if flags.I {
			pattern = `(?i)` + pattern
		}

		if flags.F {
			pattern = regexp.QuoteMeta(pattern)
		}

		reg, err := regexp.Compile(pattern)
		if err != nil {
			return
		}

		for i, v := range text {

			isContains := reg.MatchString(v)
			if flags.V {
				isContains = !isContains
			}

			if isContains {
				if flags.CC {
					count++
				} else {

					leftBorder := 0

					if flags.A != -1 {
						needWriteCount = flags.A
					}
					if flags.B != -1 {
						leftBorder = flags.B
					}
					if flags.C != -1 {
						leftBorder = flags.C
						needWriteCount = flags.C
					}

					leftBorder = i - leftBorder
					if leftBorder < 0 {
						leftBorder = 0
					}

					if leftBorder <= lastAcceptedIndex {
						leftBorder = lastAcceptedIndex + 1
					}

					if lastAcceptedIndex != -1 && leftBorder-lastAcceptedIndex > 1 {
						if flags.A != -1 || flags.B != -1 || flags.C != -1 {
							writer.WriteString("--\n")
						}
					}

					for line := leftBorder; line <= i; line++ {

						toWrite := text[line] + "\n"

						if flags.N {
							toWrite = strconv.Itoa(line) + ":" + toWrite
						}

						_, err := writer.WriteString(toWrite)
						if err != nil {
							return
						}
					}
				}
				lastAcceptedIndex = i
			} else if needWriteCount > 0 {
				toWrite := text[i] + "\n"

				if flags.N {
					toWrite = strconv.Itoa(i) + "-" + toWrite
				}

				writer.WriteString(toWrite)
				needWriteCount--
				lastAcceptedIndex = i
			}

		}
		writer.Flush()

		if flags.CC {
			fmt.Println(count)
		}
	}
}
