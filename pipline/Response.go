package pipline

type Response interface {
	SendMessage(message OutMessage)
}

type ResponseImpl struct {
	Response
	OutQueue chan OutMessage
}

func (response *ResponseImpl) SendMessage(message OutMessage) {
	response.OutQueue <- message
}

