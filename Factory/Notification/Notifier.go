package Notification

type Notifier interface {
	SendNotification(Notification) string
}