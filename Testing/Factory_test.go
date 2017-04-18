package Testing

import (
	"strings"
	"testing"

	"../Factory"
	"../Factory/Notification"
)

func TestGetSMSNotifier(t *testing.T) {
	notification := Notification.Notification{"5552233", "Hello world!"}
	notifier, err := Factory.GetNotifier(Notification.SMS)

	if err != nil {
		t.Fatal("Notifier type 'SMS' must be returned")
	}

	resultStr := notifier.SendNotification(notification)
	if !strings.Contains(resultStr, "was sent to number") {
		t.Error("SMS notifier result message was not correct.")
	}
}

func TestGetEmailNotifier(t *testing.T) {
	notification := Notification.Notification{"5552233", "Hello world!"}
	notifier, err := Factory.GetNotifier(Notification.Email)

	if err != nil {
		t.Fatal("Notifier type 'Email' must be returned")
	}

	resultStr := notifier.SendNotification(notification)
	if !strings.Contains(resultStr, "was sent to email") {
		t.Error("Email notifier result message was not correct.")
	}
}

func TestGetNonExistanceNotifier(t *testing.T) {
	_, err := Factory.GetNotifier(3)

	if err == nil {
		t.Fatal("Error must be returned for unrecognized notifier type.")
	}
}
