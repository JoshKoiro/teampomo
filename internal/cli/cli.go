package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JoshKoiro/teampomo/internal/timer"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start [duration]",
	Short: "Start a new Pomodoro session",
	Long:  `Starts a new Pomodoro session with a specified duration in minutes or seconds (e.g., "25" for 25 minutes or "15sec" for 15 seconds).`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		durationStr := "25"
		if len(args) > 0 {
			durationStr = args[0]
		}
		var duration int
		var err error

		if strings.HasSuffix(durationStr, "sec") {
			secondsStr := strings.TrimSuffix(durationStr, "sec")
			duration, err = strconv.Atoi(secondsStr)
			if err != nil {
				fmt.Println("Invalid duration. Please specify a valid number of seconds (e.g., '15sec').")
				return
			}
			fmt.Printf("Starting a new Pomodoro session for %d seconds...\n", duration)
		} else {
			durationStr = strings.TrimSpace(durationStr)
			duration, err = strconv.Atoi(durationStr)
			if err != nil {
				fmt.Println("Invalid duration. Please specify a valid number of minutes or seconds using the '15sec' syntax.")
				return
			}
			duration *= 60 // Convert minutes to seconds for consistency
			fmt.Printf("Starting a new Pomodoro session for %d minutes...\n", duration/60)
		}

		pomodoroTimer := timer.NewTimer(duration)
		pomodoroTimer.Start()
	},
}
