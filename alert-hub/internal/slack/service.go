package slack

import "fmt"

type Client struct {
	SlackId string
}

func (s Client) Send(msg string) {
	fmt.Printf("\nSending \"%s\" message to \"%s\", via slack\n", msg, s.SlackId)
}
