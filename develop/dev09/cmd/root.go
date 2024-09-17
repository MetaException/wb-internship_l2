package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wget [url]",
	Short: "a non-interactive network retriever.",
	Args:  cobra.ExactArgs(1),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		url := args[0]

		resp, err := http.Get("http://" + url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		fmt.Println(resp)

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		file, err := os.Create("received.txt") //Можно обработать content-type и сохранять файл в нужное расширение
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = file.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
