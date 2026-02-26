package sms

import "fmt"

type Client struct {
	PhoneNumber string
}

func (s Client) Send(msg string) {
	fmt.Printf("\nSending message : \"%s\", to \"%s\" via SMS \n", msg, s.PhoneNumber)
}
