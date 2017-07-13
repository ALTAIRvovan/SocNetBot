package pipline

type Executor struct {
	RequestQueue chan InMessage
	ResponseQueue chan OutMessage
}

func (ececutor *Executor) Run() {
	var response = &ResponseImpl{OutQueue: ececutor.ResponseQueue}
	for request := range ececutor.RequestQueue {
		command := request.GetCmd()
		command.Execute(request, response)
	}
}