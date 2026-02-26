package email

import "fmt"

type Client struct {
	Email string
}

func (e Client) Send(msg string) {
	fmt.Printf("\nSending message \"%s\", to \"%s\" via email\n", msg, e.Email)
}
