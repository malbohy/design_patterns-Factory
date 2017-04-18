package Factory

import (
	"errors"
	"fmt"
	"./Notification"
)

// Factory .....
func GetNotifier(notifierType int) (Notification.Notifier, error) {
	switch notifierType {
	case Notification.SMS:
		return new(Notification.SMSNotification), nil
	case Notification.Email:
		return new(Notification.EmailNotification), nil
	default:
		return nil, errors.New(fmt.Sprintf("Notifier.go type %d not recognized.", notifierType))
	}
}
