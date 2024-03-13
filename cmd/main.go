package main

import (
	"fmt"
	"os"

	"github.com/JoshKoiro/teampomo/internal/cli"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gopomodoro",
		Short: "GoPomodoro is a CLI Pomodoro timer",
		Long: `A simple and efficient CLI Pomodoro timer built with Go.
Complete documentation is available at http://gopomodoro.example.com`,
	}

	rootCmd.AddCommand(cli.StartCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
