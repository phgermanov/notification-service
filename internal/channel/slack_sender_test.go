package channel

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlack_Send(t *testing.T) {
	tests := []struct {
		name       string
		webhookURL string
		setup      func(client http.Client)
		server     *httptest.Server
		expectErr  error
	}{
		{
			name: "Sending message to Slack returns no error",
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				rw.Write([]byte(`OK`)) // send data to our test client
			})),
			expectErr: nil,
		},
		{
			name: "Sending message to Slack returns error",
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte(`invalid_token`)) // send data to our test client
			})),
			expectErr: fmt.Errorf("request failed: invalid_token"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSlack(tt.server.URL, tt.server.Client())

			err := s.Send("Hello")
			if tt.expectErr != nil {
				assert.EqualError(t, err, tt.expectErr.Error())
			}
		})
	}
}

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
