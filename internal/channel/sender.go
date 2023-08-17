package channel

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Sender
type Sender interface {
	GetName() string
	Send(s string) error
}
