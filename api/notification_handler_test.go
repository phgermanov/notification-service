package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phgermanov/notification-service/internal/notification"
	"github.com/stretchr/testify/assert"
)

// TestSendNotificationHandler is a unit test for the SendNotificationHandler function.
func TestSendNotificationHandler(t *testing.T) {
	// Define test cases with input notifications and expected HTTP response statuses.
	tests := []struct {
		name           string
		input          SendNotificationRequest
		expectedStatus int
	}{
		{
			name: "Valid Notification",
			input: SendNotificationRequest{
				Channels: []string{"channel1", "channel2"},
				Message:  "Test message",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid Notification",
			input: SendNotificationRequest{
				Channels: []string{},
				Message:  "",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	// Iterate through the test cases and run each test.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal the input notification into JSON format for the request body.
			reqBody, _ := json.Marshal(tt.input)
			req, err := http.NewRequest("POST", "/notifications", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			// Create a new recorder to capture the HTTP response.
			rr := httptest.NewRecorder()

			// Create a new instance of the API handler for testing.
			handler := NewHandler(notification.NewNotifier(0))

			// Set up the router and send the HTTP request to the handler.
			router := SetupRouter(handler)
			router.ServeHTTP(rr, req)

			// Assert that the actual HTTP response status matches the expected status.
			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}
