package notification

import (
	"errors"
	"testing"

	"github.com/phgermanov/notification-service/internal/channel/channelfakes"
	"github.com/stretchr/testify/assert"
)

func TestSendWithRetries(t *testing.T) {
	mockName := "mock"

	// Define test cases
	tests := []struct {
		name              string
		mockSender        *channelfakes.FakeSender
		setup             func(m *channelfakes.FakeSender)
		expectedErr       error
		expectedCallCount int
	}{
		{
			name: "SuccessWithZeroRetries",
			setup: func(mockSender *channelfakes.FakeSender) {
				mockSender.GetNameReturns(mockName)
			},
			mockSender:        new(channelfakes.FakeSender),
			expectedErr:       nil,
			expectedCallCount: 1,
		},
		{
			name: "SuccessAfterOneRetry",
			setup: func(mockSender *channelfakes.FakeSender) {
				mockSender.GetNameReturns(mockName)
				mockSender.SendReturns(errors.New("some error"))
				mockSender.SendReturnsOnCall(1, nil)
			},
			mockSender:        new(channelfakes.FakeSender),
			expectedErr:       nil,
			expectedCallCount: 2,
		},
		{
			name: "FailedRetries",
			setup: func(mockSender *channelfakes.FakeSender) {
				mockSender.GetNameReturns(mockName)
				mockSender.SendReturns(errors.New("some error"))
			},
			mockSender:        new(channelfakes.FakeSender),
			expectedErr:       ErrRetryFailed,
			expectedCallCount: 3,
		},
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Notifier and NotifierWorker
			n := NewNotifier(0)
			nw := NewNotifierWorker(n, 0)

			// Perform setup if provided
			if tt.setup != nil {
				tt.setup(tt.mockSender)
			}
			n.AddChannelSender(tt.mockSender)

			// Create a test notification
			notification := Notification{
				Channel: tt.mockSender.GetName(),
				Message: "test message",
			}

			// Call trySendNotification and assert results
			err := nw.trySendNotification(notification, 3)

			// Compare actual and expected errors
			if err != tt.expectedErr {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}

			// Check the number of Send calls made on the mockSender
			assert.Equal(t, tt.mockSender.SendCallCount(), tt.expectedCallCount)
		})
	}
}

func TestHandleNotification(t *testing.T) {
	mockName := "mock"

	// Define test cases
	tests := []struct {
		name              string
		mockSender        *channelfakes.FakeSender
		setup             func(m *channelfakes.FakeSender)
		expectedErr       error
		expectedCallCount int
	}{
		{
			name: "SuccessWithZeroRetries",
			setup: func(mockSender *channelfakes.FakeSender) {
				mockSender.GetNameReturns(mockName)
			},
			mockSender:        new(channelfakes.FakeSender),
			expectedErr:       nil,
			expectedCallCount: 1,
		},
		{
			name: "SuccessAfterOneRetry",
			setup: func(mockSender *channelfakes.FakeSender) {
				mockSender.GetNameReturns(mockName)
				mockSender.SendReturns(errors.New("some error"))
				mockSender.SendReturnsOnCall(1, nil)
			},
			mockSender:        new(channelfakes.FakeSender),
			expectedErr:       nil,
			expectedCallCount: 2,
		},
		{
			name: "FailedRetries",
			setup: func(mockSender *channelfakes.FakeSender) {
				mockSender.GetNameReturns(mockName)
				mockSender.SendReturns(errors.New("some error"))
			},
			mockSender:        new(channelfakes.FakeSender),
			expectedErr:       ErrRetryFailed,
			expectedCallCount: 3,
		},
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Notifier
			n := NewNotifier(0)
			nw := NewNotifierWorker(n, 0)

			// Perform setup if provided
			if tt.setup != nil {
				tt.setup(tt.mockSender)
			}
			n.AddChannelSender(tt.mockSender)

			// Create a test notification
			notification := Notification{
				Channel: tt.mockSender.GetName(),
				Message: "test message",
			}

			// Call handleNotification and assert results
			nw.handleNotification(notification)
			assert.Equal(t, tt.mockSender.SendCallCount(), tt.expectedCallCount)
		})
	}
}

func TestSendNotification(t *testing.T) {
	mockName := "mock"

	// Define test cases
	testCases := []struct {
		name          string
		setup         func(notifier *Notifier)
		notification  Notification
		expectedError error
	}{
		{
			name: "Sends notification",
			setup: func(notifier *Notifier) {
				notifier.AddChannelSender(mockWithName(mockName))
			},
			notification: Notification{
				Channel: mockName,
				Message: "Test message",
			},
			expectedError: nil,
		},
		{
			name: "Channel not found",
			notification: Notification{
				Channel: mockName,
				Message: "Test message",
			},
			expectedError: ErrChannelNotFound,
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		// Create a new Notifier and NotifierWorker
		n := NewNotifier(0)
		nw := NewNotifierWorker(n, 0)

		t.Run(tc.name, func(t *testing.T) {
			// Perform setup if provided
			if tc.setup != nil {
				tc.setup(n)
			}

			// Call sendNotification and assert results
			err := nw.sendNotification(tc.notification)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
