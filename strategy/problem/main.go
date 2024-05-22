package main

import "fmt"

type Notification struct {
	notifierType string
}

func (n *Notification) SendNotification(message string) {
	if n.notifierType == "EMAIL" {
		fmt.Printf("Sending email: %s \n", message)
	} else if n.notifierType == "SMS" {
		fmt.Printf("Sending sms: %s \n", message)
	}
}

func main() {
	s := Notification{notifierType: "SMS"}
	s.SendNotification("Hello world")
}