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

type UserIDResponse struct {
	ID string `json:"id"`
}

func GetUserID(key string) (string, error) {
	// create get request
	req, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// add headers to request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}

	// collect response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// parse response body
	var id UserIDResponse
	err = json.Unmarshal(body, &id)
	if err != nil {
		return "", fmt.Errorf("error parsing response body: %w", err)
	}

	return id.ID, nil
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

type StatusRequestBody struct {
	StatusMessage struct {
		Message struct {
			Content     string `json:"content"`
			ContentType string `json:"contentType"`
		} `json:"message"`
	} `json:"statusMessage"`
}

func SetStatusMessage(key string, message string, expireDate string) error {
	// Assume GetUserID is a function that retrieves the user ID based on the provided key
	userID, err := GetUserID(key)
	if err != nil {
		return fmt.Errorf("error getting user ID: %w", err)
	}

	// Create post body
	postBody := StatusRequestBody{
		StatusMessage: struct {
			Message struct {
				Content     string `json:"content"`
				ContentType string `json:"contentType"`
			} `json:"message"`
		}{
			Message: struct {
				Content     string `json:"content"`
				ContentType string `json:"contentType"`
			}{
				Content:     message,
				ContentType: "text", // Assuming the ContentType is HTML, adjust as necessary
			},
		},
	}

	// Convert post body to json
	postBodyJSON, err := json.Marshal(postBody)
	if err != nil {
		return fmt.Errorf("error marshaling post body to JSON: %w", err)
	}

	// Create a new HTTP request with the JSON body
	endpoint := "https://graph.microsoft.com/v1.0/users/" + userID + "/presence/setStatusMessage"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(postBodyJSON))
	if err != nil {
		return fmt.Errorf("error creating request with JSON body: %w", err)
	}

	// Add headers to request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	// Send post request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error setting status message with status code %d: %s", resp.StatusCode, resp.Status)
	}

	return nil
}
