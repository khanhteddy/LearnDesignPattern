package main

import "LearnDesignPattern/Behavioral/Strategy/solution"

func main() {
	//s := problem.NotificationService{NotifierType: "email"}
	//s.SendNotifier("Send mail")

	emailService := solution.NewNotificationService(solution.EmailNotifier{})
	emailService.SendMessage("Hello via Email")
}
