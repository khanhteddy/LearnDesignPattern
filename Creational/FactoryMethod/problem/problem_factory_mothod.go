package problem_factory

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)\n", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
}

type Service struct {
	Notifier Notifier
}

func (s Service) SendNotification(message string) {
	s.Notifier.Send(message)
}
