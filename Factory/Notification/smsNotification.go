package Notification

import "fmt"


type SMSNotification struct{}

func (n SMSNotification) SendNotification(notification Notification) string {
	return fmt.Sprintf("`%s` was sent to number `%s` successfully!", notification.Message, notification.To)
}

