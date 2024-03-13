package teamsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PresenceResponse struct {
	Availability  string `json:"availability"`
	StatusMessage string `json:"statusMessage"`
}

type DateTime struct {
	DateTime string `json:"dateTime"`
	Timezone string `json:"timeZone"`
}

type CalEvent struct {
	Subject string   `json:"subject"`
	Start   DateTime `json:"start"`
	End     DateTime `json:"end"`
}

func CreateEvent(key string, subject string, start string, end string) error {
	// create post request
	req, err := http.NewRequest("POST", "https://graph.microsoft.com/beta/me/events", nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// add headers to request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	// create post body
	postBody := CalEvent{
		Subject: subject,
		Start:   DateTime{DateTime: start, Timezone: "UTC"},
		End:     DateTime{DateTime: end, Timezone: "UTC"},
	}

	postBody.Start.DateTime = start
	postBody.End.DateTime = end

	// convert post body to json
	postBodyJson, err := json.Marshal(postBody)
	if err != nil {
		return fmt.Errorf("error converting post body to json: %w", err)
	}
	fmt.Println(string(postBodyJson))
	// add post body to request
	req.Body = io.NopCloser(io.MultiReader(bytes.NewBuffer(postBodyJson)))

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// check status code
	if resp.StatusCode != 201 {
		return fmt.Errorf("error creating event: %s", resp.Status)
	}

	// collect response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// parse the body to return the id of the created event
	var event CalEvent

	err = json.Unmarshal(body, &event)
	if err != nil {
		return fmt.Errorf("error parsing response body: %w", err)
	}

	return nil
}

func GetStatus(key string) (PresenceResponse, error) {

	// Create request
	req, err := http.NewRequest("GET", "https://graph.microsoft.com/beta/me/presence", nil)
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers to request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	// Send request and get response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error sending request: %w", err)
	}

	// collect response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	// parse response body
	availability, statusMessage, err := getPresence(string(body))
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error parsing response body: %w", err)
	}

	// Return response body
	return PresenceResponse{Availability: availability, StatusMessage: statusMessage}, err
}

func getPresence(responseJSON string) (string, string, error) {
	var presence PresenceResponse

	err := json.Unmarshal([]byte(responseJSON), &presence)
	if err != nil {
		return "", "", fmt.Errorf("error parsing response body: %w", err)
	}

	return presence.Availability, presence.StatusMessage, err
}
