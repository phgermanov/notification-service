package notification

import (
	"runtime"
	"testing"

	"github.com/phgermanov/notification-service/internal/channel"
	"github.com/phgermanov/notification-service/internal/channel/channelfakes"
	"github.com/stretchr/testify/assert"
)

func TestAddChannelSender(t *testing.T) {
	// Create a mock channel sender
	mock := mockWithName("mock")

	// Define test cases
	testCases := []struct {
		name          string
		channel       channel.Sender
		setup         func(notifier *Notifier)
		expectedError error
	}{
		{
			name:          "Add a channel",
			channel:       mock,
			expectedError: nil,
		},
		{
			name: "Add a channel when one already exists",
			setup: func(notifier *Notifier) {
				notifier.AddChannelSender(mock)
			},
			channel:       mock,
			expectedError: ErrAlreadyExists,
		},
	}

	// Create a new notifier instance
	notifier := NewNotifier(0)

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform setup if provided
			if tc.setup != nil {
				tc.setup(notifier)
			}

			// Attempt to add the channel sender and assert the result
			err := notifier.AddChannelSender(tc.channel)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetChannels(t *testing.T) {
	// Create mock channel senders
	mock1 := mockWithName("mock1")
	mock2 := mockWithName("mock2")

	// Define test cases
	testCases := []struct {
		name          string
		channels      []channel.Sender
		expectedNames []string
	}{

		{
			name:          "Get all channels",
			channels:      []channel.Sender{mock1, mock2},
			expectedNames: []string{mock1.GetName(), mock2.GetName()},
		},
	}

	// Create a new notifier instance
	notifier := NewNotifier(0)

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Add mock channels to the notifier
			for _, ch := range tc.channels {
				_ = notifier.AddChannelSender(ch)
			}

			// Get the list of channel names and assert against the expected names
			names := notifier.GetChannels()
			assert.ElementsMatch(t, tc.expectedNames, names)
		})
	}
}

func TestGetChannel(t *testing.T) {
	// Create a mock channel sender
	mock := mockWithName("mock")

	// Define test cases
	testCases := []struct {
		name            string
		channels        []channel.Sender
		getName         string
		expectedChannel channel.Sender
		expectedError   error
	}{
		{
			name:            "Get a channel",
			channels:        []channel.Sender{mock},
			getName:         mock.GetName(),
			expectedChannel: mock,
			expectedError:   nil,
		},
		{
			name:            "Get non-existing channel",
			channels:        []channel.Sender{},
			getName:         mock.GetName(),
			expectedChannel: nil,
			expectedError:   ErrChannelNotFound,
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		// Create a new notifier instance
		notifier := NewNotifier(0)
		t.Run(tc.name, func(t *testing.T) {
			// Add mock channels to the notifier
			for _, ch := range tc.channels {
				_ = notifier.AddChannelSender(ch)
			}

			// Get the channel sender by name and assert against the expected results
			ch, err := notifier.getChannelSender(tc.getName)
			assert.Equal(t, tc.expectedChannel, ch)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestStartWorkers(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name       string
		numWorkers int
	}{
		{
			name:       "1 worker",
			numWorkers: 1,
		},
		{
			name:       "5 workers",
			numWorkers: 5,
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Notifier
			notifier := NewNotifier(0)

			// Get the number of goroutines before starting workers
			numGoroutinesBefore := runtime.NumGoroutine()

			// Start workers
			notifier.StartWorkers(tc.numWorkers)

			// Get the number of goroutines after starting workers
			numGoroutinesAfter := runtime.NumGoroutine()

			// Check if the number of goroutines increased by the expected amount
			assert.Equal(t, numGoroutinesBefore+tc.numWorkers, numGoroutinesAfter)
		})
	}
}

func TestEnqueueNotifications(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name          string
		notifications []Notification
	}{
		{
			name: "1 notification",
			notifications: []Notification{
				{
					Channel: "foo",
					Message: "hello",
				},
			},
		},
		{
			name: "3 notifications",
			notifications: []Notification{
				{
					Channel: "foo",
					Message: "hello",
				},
				{
					Channel: "bar",
					Message: "hello",
				},
				{
					Channel: "baz",
					Message: "hello",
				},
			},
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Notifier
			notifier := NewNotifier(0)

			// Enqueue notifications
			notifier.EnqueueNotifications(tc.notifications)

			// Check if the length of the notification pool increased by the expected amount
			assert.Equal(t, len(tc.notifications), len(notifier.notificationPool))
		})
	}
}

func mockWithName(name string) channel.Sender {
	// Create and return a mock channel sender with the given name
	mock := new(channelfakes.FakeSender)
	mock.GetNameReturns(name)
	return mock
}
