package pipline

type OutMessageSender func(OutMessage)

type Sender struct {
	addresses map[string]OutMessageSender
	messageQueue chan OutMessage
}

func NewSender(messageQueue chan OutMessage) *Sender {
	sender := new(Sender)
	sender.addresses = make(map[string]OutMessageSender)
	sender.messageQueue = messageQueue
	return sender
}

func (sender *Sender)AddSender(name string, messageSender OutMessageSender) {
	sender.addresses[name] = messageSender
}

func (sender *Sender)send(message OutMessage) {
	handler, ok := sender.addresses[message.GetType()]
	if !ok {
		panic("Sender for this type was not found")
	}
	handler(message)
}

func (sender *Sender)Run() {
	for message := range sender.messageQueue {
		sender.send(message)
	}
}
