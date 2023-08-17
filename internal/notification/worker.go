package notification

import (
	"errors"
	"log"
	"time"
)

var (
	ErrRetryFailed     = errors.New("failed to send notification after multiple retries")
	ErrChannelNotFound = errors.New("channel not found")
)

// Worker is an interface that defines the behavior of a notification worker.
type Worker interface {
	Start()
}

// NotifierWorker represents a worker responsible for processing notifications.
type NotifierWorker struct {
	*Notifier
	WorkerID int
}

// NewNotifierWorker creates a new NotifierWorker instance.
func NewNotifierWorker(notifier *Notifier, workerID int) *NotifierWorker {
	return &NotifierWorker{Notifier: notifier, WorkerID: workerID}
}

// Start starts the worker to process notifications.
func (w *NotifierWorker) Start() {
	defer w.wg.Done()
	for notification := range w.notificationPool {
		w.handleNotification(notification)
	}
}

// handleNotification processes a single notification.
func (w *NotifierWorker) handleNotification(notification Notification) {
	log.Printf("worker %d processing notification: %+v\n", w.WorkerID, notification)
	if err := w.trySendNotification(notification, 3); err != nil {
		log.Printf("error: Failed to send notification after multiple tries: %v\n", err)
	}
}

// trySendNotification tries to send a notification multiple times with retries.
func (w *NotifierWorker) trySendNotification(notification Notification, maxRetries int) error {
	for retries := 0; retries < maxRetries; retries++ {
		if err := w.sendNotification(notification); err == nil {
			return nil
		}
		log.Printf("retrying in %v\n", w.retryDuration)
		time.Sleep(w.retryDuration)
	}
	return ErrRetryFailed
}

// sendNotification sends the notification to the specified channels.
func (w *NotifierWorker) sendNotification(notification Notification) error {
	channelSender, err := w.getChannelSender(notification.Channel)
	if err != nil {
		return err
	}
	if err := channelSender.Send(notification.Message); err != nil {
		return err
	}
	return nil
}
