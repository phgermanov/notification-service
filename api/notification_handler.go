package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phgermanov/notification-service/internal/notification"
)

// Define a custom error for invalid requests.
var (
	ErrInvalidRequest = errors.New("invalid request")
)

// Notification represents a message to be sent to multiple channels.
type SendNotificationRequest struct {
	Channels []string `json:"channels" binding:"required"`
	Message  string   `json:"message" binding:"required"`
}

// SendNotificationHandler handles the HTTP request for sending notifications.
func (h Handler) SendNotificationHandler(c *gin.Context) {
	// Create a variable to hold the incoming notification data.
	var input SendNotificationRequest

	// Attempt to bind the incoming JSON data to the input variable.
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Enqueue the received notification for processing in the queue.
	h.notifier.EnqueueNotifications(mapInputToNotification(input))

	// Respond with a success message indicating that the notification was accepted for processing.
	c.JSON(http.StatusOK, gin.H{"message": "Notification accepted for processing"})
}

func mapInputToNotification(input SendNotificationRequest) (notifications []notification.Notification) {
	for _, channel := range input.Channels {
		notifications = append(notifications, notification.Notification{
			Channel: channel,
			Message: input.Message,
		})
	}
	return
}
