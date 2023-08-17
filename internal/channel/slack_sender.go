package channel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Slack represents a Slack channel for sending messages.
type Slack struct {
	webhookURL string       // The webhook URL for sending messages to Slack.
	name       string       // The name of the sender.
	client     *http.Client // HTTP client for making requests.
}

// PostMessageRequest is the structure used to define the JSON request for posting a message to Slack.
type PostMessageRequest struct {
	Text string `json:"text"` // The text content of the message.
}

// NewSlack creates a new Slack channel instance.
func NewSlack(webhookURL string, client *http.Client) *Slack {
	return &Slack{
		webhookURL: webhookURL,
		name:       "Slack",
		client:     client,
	}
}

// Send sends a message to the Slack channel.
func (s *Slack) Send(message string) error {
	// Create a JSON request body.
	reqBody, err := json.Marshal(PostMessageRequest{
		Text: message,
	})
	if err != nil {
		return err
	}

	// Send the POST request to the Slack webhook URL.
	resp, err := s.client.Post(s.webhookURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("request failed: %s", string(b))
	}
	log.Printf("slack message sent: %s", message)
	defer resp.Body.Close() // Close the response body when done with it.

	return nil
}

// GetName returns the name of the Slack sender.
func (s *Slack) GetName() string {
	return s.name
}
