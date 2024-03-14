package timer

import (
	"fmt"
	"time"

	"github.com/JoshKoiro/teampomo/internal/teamsapi"
	"github.com/fatih/color"
)

type Timer struct {
	Duration time.Duration
}

// NewTimer now expects duration in seconds for flexibility.
func NewTimer(seconds ...int) *Timer {
	duration := 25 * 60 // 25 minutes default
	if len(seconds) > 0 {
		duration = seconds[0]
	}
	return &Timer{
		Duration: time.Second * time.Duration(duration),
	}
}

func (t *Timer) Start(key string, duration int) {
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			remainingTime := endTime.Sub(time.Now()).Truncate(time.Second)
			minutes := remainingTime / time.Minute
			seconds := remainingTime % time.Minute / time.Second
			if time.Now().After(endTime) {
				// Clear the last line from console output
				fmt.Printf("\033[2K\r")
				c := color.New(color.FgHiGreen)
				c.Println("\n🍅Pomodoro completed!🍅\n")
				return
			} else {
				// Update the console output
				c := color.New(color.FgHiCyan)
				c.Printf("\033[2K\rTime remaining: %d m %d sec", minutes, seconds)
			}

			// Update the teams status message every minute
			if time.Now().Second()%60 == 0 {
				// If duration/60 is 1 or less minutes, update the status message to say minute left instead of minutes left
				if minutes == 1 {
					teamsError := teamsapi.SetStatusMessage(key, "Busy in a Pomodoro session for the next "+fmt.Sprintf("%d minute...🍅", minutes), "")
					if teamsError != nil {
						fmt.Println(teamsError)
					}
				} else if minutes < 1 {
					teamsError := teamsapi.SetStatusMessage(key, "Less than a minute left in a Pomodoro session...🍅", "")
					if teamsError != nil {
						fmt.Println(teamsError)
					}
				} else {
					teamsError := teamsapi.SetStatusMessage(key, "Busy in a Pomodoro session for the next "+fmt.Sprintf("%d minutes...🍅", minutes), "")
					if teamsError != nil {
						fmt.Println(teamsError)
					}
				}

			}
		}
	}
}
