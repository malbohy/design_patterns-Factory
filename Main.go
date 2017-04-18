package main
import (
	"./Factory"
	"./Factory/Notification"
	"fmt"
)

func main(){
	fmt.Println("moahemd")
	message , _ := Factory.GetNotifier(1)
	fmt.Println(message)
	//
	notification := Notification.Notification{"mohamed", "says heelo from main"}
	messagerecived :=  message.SendNotification(notification)
	//
	fmt.Println(messagerecived)


}