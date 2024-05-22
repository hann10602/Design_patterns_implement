package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {}

type SMSNotifier struct {}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending email: %s \n", message)
}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending sms: %s \n", message)
}

type NotificationService struct {
	notifier Notifier
}

func (n NotificationService) SendNotification(message string) {
	n.notifier.Send(message)
}

func main() {
	s := NotificationService{
		notifier: SMSNotifier{},
	}

	s.SendNotification("Hello world")
}