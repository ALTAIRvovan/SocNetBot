package pipline

type Request interface {
	GetCmd() *Command
	GetMessage() InMessage
}