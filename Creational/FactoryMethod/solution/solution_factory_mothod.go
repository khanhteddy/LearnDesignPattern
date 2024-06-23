package solution_factory

import "fmt"

type Notifier interface {
	Send(message string)
}

type emailNotifier struct{}

func (emailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)", message)
}

type sMSNotifier struct{}

func (sMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)", message)
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(message string) {
	s.notifier.Send(message)
}

func NewNotificationService(notifier Notifier) *Service {
	return &Service{notifier: notifier}
}

func CreateNotifier(t string) Notifier {
	if t == "sms" {
		return sMSNotifier{}
	}

	return emailNotifier{}
}
