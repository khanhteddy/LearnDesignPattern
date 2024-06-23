package solution

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Send message %s (Send Email)\n", message)
}

type SmsNotifier struct {
}

func (SmsNotifier) Send(message string) {
	fmt.Printf("Send message %s (Send Email)\n", message)
}

type NotificationService struct {
	notifier Notifier
}

//	func (s NotificationService) SendMessage(message string) {
//		s.notifier.Send(message)
//	}
func NewNotificationService(notifier Notifier) *NotificationService {
	return &NotificationService{notifier: notifier}
}

func (s *NotificationService) SendMessage(message string) {
	s.notifier.Send(message)
}
