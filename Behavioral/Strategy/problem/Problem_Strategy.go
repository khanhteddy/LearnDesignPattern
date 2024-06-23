package problem

import "fmt"

type NotificationService struct {
	NotifierType string
}

func (s NotificationService) SendNotifier(message string) {
	if s.NotifierType == "email" {
		fmt.Printf("Send message %s (Send Email)\n", message)
	} else if s.NotifierType == "sms" {
		fmt.Printf("Send message %s (Send sms)\n", message)
	}
}
