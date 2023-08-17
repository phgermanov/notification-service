package api

import (
	"github.com/phgermanov/notification-service/internal/notification"
)

type Handler struct {
	notifier *notification.Notifier
}

func NewHandler(notifier *notification.Notifier) *Handler {
	return &Handler{notifier}
}
