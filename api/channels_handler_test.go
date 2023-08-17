package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phgermanov/notification-service/internal/notification"
	"github.com/stretchr/testify/assert"
)

// TestGetChannelsHandler is a unit test for the GetChannels function.
func TestGetChannelsHandler(t *testing.T) {
	// Define test cases with expected HTTP response statuses.
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Valid Notification",
			expectedStatus: http.StatusOK,
			// TODO: Check body if needed for this test case.
		},
	}

	// Iterate through the test cases and run each test.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request for retrieving channels.
			req, err := http.NewRequest("GET", "/channels", nil)
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
			// TODO: Add assertions for the response body if needed.
		})
	}
}
