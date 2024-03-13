package timer

import (
	"fmt"
	"time"
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
			//format the time remaining and return it to the console.
			//clear only the last printed line from console output
			//format the time remaining as a clock with minutes and seconds with a : in between them
			fmt.Printf("\033[2K\rTime remaining: %v", endTime.Sub(time.Now()).Truncate(time.Second))
		}
	}
}
