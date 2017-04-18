package main

import (
	"fmt"

	"./Factory"
	"./Factory/Notification"
)

func main() {
	sendAnSMSNotification()
	sendAnEmailNotification()
}

func sendAnSMSNotification() {
	SMSmessage, errorResult := Factory.GetNotifier(Notification.SMS)
	mobileNumber := "+201114283495"
	SMSContent := "This Is An Sms Content"
	if errorResult == nil {
		SMSnotification := Notification.Notification{mobileNumber, SMSContent}
		messagerecived := SMSmessage.SendNotification(SMSnotification)
		fmt.Println(messagerecived)
	} else {
		fmt.Println("there is an error happened", errorResult)
	}
}

func sendAnEmailNotification() {
	emailmessage, errorResult := Factory.GetNotifier(Notification.Email)
	email := "malbohy@gmail.com"
	emailContent := "This Is An Email Content"
	if errorResult == nil {
		emailNotification := Notification.Notification{email, emailContent}
		messagerecived := emailmessage.SendNotification(emailNotification)
		fmt.Println(messagerecived)
	} else {
		fmt.Println("there is an error happened", errorResult)
	}
}
