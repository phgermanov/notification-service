package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetChannels handles the HTTP request for retrieving the list of channels.
func (h Handler) GetChannels(c *gin.Context) {
	// Retrieve the list of channels from the notifier.
	channels := h.notifier.GetChannels()

	// Respond with the list of channels in JSON format.
	c.JSON(http.StatusOK, gin.H{"channels": channels})
}
