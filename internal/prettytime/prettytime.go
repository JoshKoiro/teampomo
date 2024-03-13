package prettytime

import (
	"time"
)

// prettyTime takes a UTC time string and converts it to a pretty local time string.
func PrettyValue(utcStr string) (string, error) {
	// Parse the input time string as UTC
	utcTime, err := time.Parse(time.RFC3339, utcStr)
	if err != nil {
		return "", err
	}

	// Convert UTC time to local time
	localTime := utcTime.Local()

	// Format the local time to a pretty string (e.g., "3:04pm")
	prettyTime := localTime.Format("3:04pm")

	return prettyTime, nil
}
