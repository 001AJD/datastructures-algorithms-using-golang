package notifier

type Notifier interface {
	Send(message string)
}

func SendAlert(n Notifier, message string) {
	n.Send(message)
}
