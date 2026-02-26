package main

import (
	"alert-hub/internal/email"
	"alert-hub/internal/notifier"
	"alert-hub/internal/slack"
	"alert-hub/internal/sms"
	"fmt"
)

func main() {
	fmt.Println("Interfaces in golang")
	client := []notifier.Notifier{
		slack.Client{
			SlackId: "A12345",
		},
		email.Client{Email: "ajinkya@gmail.com"},
		sms.Client{PhoneNumber: "123456789"},
	}

	for _, v := range client {
		v.Send("System failure in production")
	}
}
