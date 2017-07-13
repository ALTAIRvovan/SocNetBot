package command

type Message interface {
	GetText() string
	SendResponse(response string)
}
