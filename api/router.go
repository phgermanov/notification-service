package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/notifications", handler.SendNotificationHandler)
	r.GET("/channels", handler.GetChannels)

	return r
}
