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
	endTime := time.Now().Add(t.Duration)

	for range time.Tick(time.Second) {
		if time.Now().After(endTime) {
			//clear the last line from console output
			fmt.Printf("\033[2K\r")
			c := color.New(color.FgHiGreen)
			c.Println("\n🍅Pomodoro completed!🍅\n")
			break
		} else {
			remainingTime := endTime.Sub(time.Now()).Truncate(time.Second)
			minutes := remainingTime / time.Minute
			seconds := remainingTime % time.Minute / time.Second

			c := color.New(color.FgHiCyan)
			c.Printf("\033[2K\rTime remaining: %d m %d sec", minutes, seconds)
		}
	}

	// update the teams status message every minute
	for range time.Tick(time.Minute) {
		// if duration/60 is 1 or less minutes, update the status message to say minute left instead of minutes left
		if duration/60 <= 1 {
			teamsError := teamsapi.SetStatusMessage(key, fmt.Sprintf("Pomodoro: %d minute left", duration/60), "")
			if teamsError != nil {
				fmt.Println(teamsError)
			}
		}
		teamsError := teamsapi.SetStatusMessage(key, fmt.Sprintf("Pomodoro: %d minutes left", duration/60), "")
		if teamsError != nil {
			fmt.Println(teamsError)
		}
	}
}
