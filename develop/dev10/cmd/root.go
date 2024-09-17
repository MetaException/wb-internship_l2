package cmd

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-telnet [flags] [host] [port]",
	Short: "telnet client",
	Args:  cobra.ExactArgs(2),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("timeout", "t", "10", "timeout (default 10s)")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		host := args[0]
		port := args[1]
		timeoutStr, err := rootCmd.Flags().GetString("timeout")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		timeout, err := ParseTimeout(timeoutStr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if timeout == 0 {
			timeout = time.Duration(time.Second * 10)
		}

		conn, err := Connect(host, port, timeout)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		doneCh := make(chan os.Signal, 1)
		signal.Notify(doneCh, syscall.SIGINT, syscall.SIGTERM)

		go Send(conn, doneCh)
		go Recieve(conn, doneCh)

		<-doneCh
		conn.Close()
	}
}

func ParseTimeout(timeoutstr string) (time.Duration, error) {
	timeMap := map[string]time.Duration{
		"s": time.Second,
		"h": time.Hour,
		"m": time.Minute,
	}

	timeout, err := strconv.Atoi(timeoutstr[:len(timeoutstr)-1])
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return time.Duration(timeMap[timeoutstr[len(timeoutstr)-1:]] * time.Duration(timeout)), nil
}

func Connect(host string, port string, timeout time.Duration) (net.Conn, error) {

	addr := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Send(conn net.Conn, doneCh chan os.Signal) {

	_, err := io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
		doneCh <- syscall.SIGQUIT
		return
	}
}

func Recieve(conn net.Conn, doneCh chan os.Signal) {

	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println(err.Error())
		doneCh <- syscall.SIGQUIT
		return
	}
}
