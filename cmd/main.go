package main

import (
	"log"
	"net/http"

	"github.com/phgermanov/notification-service/api"
	"github.com/phgermanov/notification-service/config"
	"github.com/phgermanov/notification-service/internal/channel"
	"github.com/phgermanov/notification-service/internal/notification"
)

func main() {
	// Load the configuration
	appConfig := config.LoadConfig()

	// Initialize the notification service
	notifier := initializeNotifier(appConfig)

	// Initialize the API service
	initializeAPI(notifier, appConfig)
}

// initializeNotifier sets up the notification service
func initializeNotifier(config config.Settings) *notification.Notifier {
	// Create a new notifier instance with the configured retry duration
	notifier := notification.NewNotifier(config.RetryDuration)

	// Start a specified number of worker goroutines for processing notifications
	notifier.StartWorkers(5)

	// Add a Slack channel sender to the notifier
	err := notifier.AddChannelSender(channel.NewSlack(config.SlackWebhookURL, http.DefaultClient))
	if err != nil {
		log.Printf("error adding Slack channel: %v", err)
	}

	// Add an Email channel sender to the notifier
	err = notifier.AddChannelSender(channel.NewEmail("example@gmail.com"))
	if err != nil {
		log.Printf("error adding Email channel: %v", err)
	}

	return notifier
}

// initializeAPI sets up the API service
func initializeAPI(notifier *notification.Notifier, config config.Settings) {
	// Create a new API handler using the provided notifier
	handler := api.NewHandler(notifier)

	// Set up the API router
	r := api.SetupRouter(handler)

	// Start the API server on the configured port
	port := config.Port
	if err := r.Run(":" + port); err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}
}
