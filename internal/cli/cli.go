package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
		// have the user paste the graph API key prior to execution
		key := "eyJ0eXAiOiJKV1QiLCJub25jZSI6IlpxMThGZ1RpczIxeXZ4SVl6eDhmTGltT2s3Z01qMWd4X24tOWVuMXNnd3ciLCJhbGciOiJSUzI1NiIsIng1dCI6IlhSdmtvOFA3QTNVYVdTblU3Yk05blQwTWpoQSIsImtpZCI6IlhSdmtvOFA3QTNVYVdTblU3Yk05blQwTWpoQSJ9.eyJhdWQiOiIwMDAwMDAwMy0wMDAwLTAwMDAtYzAwMC0wMDAwMDAwMDAwMDAiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC9iYTJiMDVmZS0yNjIwLTQ4YTctYmRjZC05NTJmYWE5NTI4MTkvIiwiaWF0IjoxNzEwMjkzNjM4LCJuYmYiOjE3MTAyOTM2MzgsImV4cCI6MTcxMDM4MDMzOCwiYWNjdCI6MCwiYWNyIjoiMSIsImFpbyI6IkFWUUFxLzhXQUFBQTU1eDNQOEEvcmpDbm9XT0VFREdYKzkrbFI1Mlc0NHRyNExBVFN4WlJMSXlPOFhsNEp1cFZZR0YvZlRFeDdjL2VNQVcvRXJvM2E2R2lxU25JckFYZ3ZJY2E3Y0Z4RkNzM1ppakpSdFg3ZS9zPSIsImFtciI6WyJwd2QiLCJtZmEiXSwiYXBwX2Rpc3BsYXluYW1lIjoiR3JhcGggRXhwbG9yZXIiLCJhcHBpZCI6ImRlOGJjOGI1LWQ5ZjktNDhiMS1hOGFkLWI3NDhkYTcyNTA2NCIsImFwcGlkYWNyIjoiMCIsImZhbWlseV9uYW1lIjoiS29pcm8iLCJnaXZlbl9uYW1lIjoiSm9zaCIsImlkdHlwIjoidXNlciIsImlwYWRkciI6IjE3NC40OS4xMzcuMTk5IiwibmFtZSI6Ikpvc2ggS29pcm8iLCJvaWQiOiIwYTExOTk5ZS1jYzQ3LTQ1NzctODc4OC0yZjlhMTZjY2JlOTMiLCJvbnByZW1fc2lkIjoiUy0xLTUtMjEtMTE3NzIzODkxNS0zMjkwNjgxNTItMTQxNzAwMTMzMy03MTQ4IiwicGxhdGYiOiIzIiwicHVpZCI6IjEwMDMwMDAwOUNGNjkyNjMiLCJyaCI6IjAuQVM0QV9nVXJ1aUFtcDBpOXpaVXZxcFVvR1FNQUFBQUFBQUFBd0FBQUFBQUFBQUF1QUZ3LiIsInNjcCI6IkNvbnRhY3RzLlJlYWRXcml0ZSBNYWlsLlJlYWQgb3BlbmlkIFByZXNlbmNlLlJlYWQuQWxsIHByb2ZpbGUgVXNlci5SZWFkIFVzZXIuUmVhZEJhc2ljLkFsbCBVc2VyQWN0aXZpdHkuUmVhZFdyaXRlLkNyZWF0ZWRCeUFwcCBlbWFpbCBDYWxlbmRhcnMuUmVhZFdyaXRlIiwic2lnbmluX3N0YXRlIjpbImttc2kiXSwic3ViIjoiYXhjOU9Hc00yR2NpSmdDSUdCR1NxUFdxakdnSTdQR3ZTcWNpT1R4c2hJbyIsInRlbmFudF9yZWdpb25fc2NvcGUiOiJOQSIsInRpZCI6ImJhMmIwNWZlLTI2MjAtNDhhNy1iZGNkLTk1MmZhYTk1MjgxOSIsInVuaXF1ZV9uYW1lIjoiSktvaXJvQHJ2aWkuY29tIiwidXBuIjoiSktvaXJvQHJ2aWkuY29tIiwidXRpIjoicEl2R2tsTFlEMHkxVzMtSjlBcEVBQSIsInZlciI6IjEuMCIsIndpZHMiOlsiYjc5ZmJmNGQtM2VmOS00Njg5LTgxNDMtNzZiMTk0ZTg1NTA5Il0sInhtc19jYyI6WyJDUDEiXSwieG1zX3NzbSI6IjEiLCJ4bXNfc3QiOnsic3ViIjoiV3k5VTA5YXBWaXZDdnFES0NSSnpRQi0yV0JvWGZXbTBSS0NUb20zeTVqVSJ9LCJ4bXNfdGNkdCI6MTUxMDkyNjI4N30.DLvs5EHBBZ4AyMyUozCyVadyiskaiqxSyz35_aVW2QkZ3VHl7igQwZ0bD7poI8YUSmc8QXwoI4Rp1adksPrOZYEMhXMSKeGu0NyI9My4w4iLEKwRFALXP3oWQkAw09IaNlMSam8Xl3cd4DklfAAwW7dlXp9UvmGOzYRyl9U8xS6jvt4pQxy7ItMxcFRPI63Qmi6byLSvOrz2qD73k5A858z9KDoMPrd6j0jYuAWyhjHq2Q-4ASklWZGi4qr6J-eIEgaFRYg2bgSNXD_1CpAflX5LGmwpsxVRDbc9J8CIDHpHovSOoo8mLAunDKv-sevNLj7r7taKdOxQnVE34KjjKA"

		// TESTING - GET RID OF THIS LATER

		currentPresence, teamsError := teamsapi.GetStatus(key)
		if teamsError != nil {
			fmt.Println(teamsError)
		}

		fmt.Println(currentPresence)

		// KEEP THIS AND BELOW...
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

		}

		pomodoroTimer := timer.NewTimer(duration)
		pomodoroTimer.Start()
	},
}
