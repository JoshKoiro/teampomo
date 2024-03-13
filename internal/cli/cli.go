package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/JoshKoiro/teampomo/internal/loadkey"
	"github.com/JoshKoiro/teampomo/internal/teamsapi"
	"github.com/JoshKoiro/teampomo/internal/timer"
	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start [duration]",
	Short: "Start a new Pomodoro session",
	Long:  `Starts a new Pomodoro session with a specified duration in minutes or seconds (e.g., "25" for 25 minutes or "15sec" for 15 seconds).`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// load the API key from a text file
		key, keyerror := loadkey.LoadKey()
		if keyerror != nil {
			fmt.Println("Error loading API key:", keyerror)
		}
		startTime := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
		var endTime string
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
			c := color.New(color.FgHiMagenta)
			c.Printf("\nStarting a new Pomodoro session for %d seconds...\n\n", duration)

			// if less than 1 min, do not create a calender event
		} else {
			durationStr = strings.TrimSpace(durationStr)
			duration, err = strconv.Atoi(durationStr)
			if err != nil {
				fmt.Println("Invalid duration. Please specify a valid number of minutes or seconds using the '15sec' syntax.")
				return
			}
			duration *= 60 // Convert minutes to seconds for consistency
			c := color.New(color.FgHiMagenta)
			c.Printf("Starting a new Pomodoro session for %d minutes...\n", duration/60)
			endTime = time.Now().Add(time.Duration(duration) * time.Second).UTC().Format("2006-01-02T15:04:05.999Z")

			// print start time
			c.Printf("Start time: %s\n", startTime)
			// print end time
			c.Printf("End time: %s\n", endTime)

			// create calendar event
			teamsError := teamsapi.CreateEvent(key, "Pomodoro", startTime, endTime)
			if teamsError != nil {
				fmt.Println(teamsError)
			}

			// set initial status message
			// teamsError = teamsapi.SetStatusMessage(key, "In a Pomodoro session for the next "+strconv.Itoa(duration/60)+" minutes", endTime)
			// if teamsError != nil {
			// 	fmt.Println(teamsError)
			// }
		}

		pomodoroTimer := timer.NewTimer(duration)
		pomodoroTimer.Start(key, duration)
	},
}
