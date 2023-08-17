package notification

import (
	"errors"
	"sync"
	"time"

	"github.com/phgermanov/notification-service/internal/channel"
)

var (
	ErrAlreadyExists = errors.New("channel already exists")
)

// Notification represents a message to be sent to multiple channels.
type Notification struct {
	Channel string
	Message string
}

// Notifier manages the sending of notifications to different channels.
type Notifier struct {
	channelSenders   map[string]channel.Sender
	notificationPool chan Notification
	wg               sync.WaitGroup
	retryDuration    time.Duration
}

// NewNotifier creates a new Notifier instance.
func NewNotifier(retryDuration time.Duration) *Notifier {
	return &Notifier{
		channelSenders:   make(map[string]channel.Sender),
		notificationPool: make(chan Notification, 100),
		retryDuration:    retryDuration,
	}
}

// StartWorkers starts a specified number of worker goroutines for processing notifications.
func (n *Notifier) StartWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		worker := NewNotifierWorker(n, i)
		go worker.Start()
		n.wg.Add(1)
	}
}

// EnqueueNotification adds a notification to the processing queue.
func (n *Notifier) EnqueueNotifications(notifications []Notification) {
	for _, notification := range notifications {
		n.notificationPool <- notification
	}
}

// AddChannelSender adds a channel sender to the Notifier.
func (n *Notifier) AddChannelSender(channelSender channel.Sender) error {
	if _, exists := n.channelSenders[channelSender.GetName()]; exists {
		return ErrAlreadyExists
	}
	n.channelSenders[channelSender.GetName()] = channelSender
	return nil
}

// GetChannels returns a list of available channel names.
func (n *Notifier) GetChannels() []string {
	keys := make([]string, 0, len(n.channelSenders))
	for k := range n.channelSenders {
		keys = append(keys, k)
	}
	return keys
}

// getChannelSender retrieves a channel sender by its name.
func (n *Notifier) getChannelSender(name string) (channel.Sender, error) {
	channelSender, found := n.channelSenders[name]
	if !found {
		return nil, ErrChannelNotFound
	}
	return channelSender, nil
}
