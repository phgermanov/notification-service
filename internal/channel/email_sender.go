package channel

import "log"

type Email struct {
	to   string
	name string
}

func NewEmail(to string) *Email {
	return &Email{
		to:   to,
		name: "Email",
	}
}

func (e *Email) Send(message string) error {
	log.Printf("sending email to: %s. Message: %s\n", e.to, message)
	// Simulate email sending process here
	return nil
}

func (e *Email) GetName() string {
	return e.name
}
