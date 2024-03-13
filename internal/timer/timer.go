package timer

import (
	"fmt"
	"time"

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

func (t *Timer) Start() {
	endTime := time.Now().Add(t.Duration)

	for range time.Tick(time.Second) {
		if time.Now().After(endTime) {
			//clear the last line from console output
			fmt.Printf("\033[2K\r")
			fmt.Println("Pomodoro completed!")
			break
		} else {
			remainingTime := endTime.Sub(time.Now()).Truncate(time.Second)
			minutes := remainingTime / time.Minute
			seconds := remainingTime % time.Minute / time.Second

			c := color.New(color.FgHiGreen)
			c.Printf("\033[2K\rTime remaining: %d m %d sec", minutes, seconds)
		}
	}
}
