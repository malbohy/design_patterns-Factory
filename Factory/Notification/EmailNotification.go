package Notification

import "fmt"

type EmailNotification struct{}

func (n EmailNotification) SendNotification(notification Notification) string {
	return fmt.Sprintf("`%s` was sent to email `%s` successfully!", notification.Message, notification.To)
}