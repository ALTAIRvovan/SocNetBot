package pipline

type Typed interface {
	GetType() string
}

type InMessage interface {
	Typed
	MakeResponse() OutMessage
	GetCmd() Command
	GetContentText() string
}

type OutMessage interface {
	Typed
	SetText(text string)
}

type MessageChannel interface {
	SetOutQueue(queue chan InMessage)
	SetCommandProvider(provider *CommandProvider)
}