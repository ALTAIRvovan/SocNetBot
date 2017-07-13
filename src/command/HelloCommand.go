package command

import "pipline"

type HelloCommand struct {
	pipline.Command
}

func (cmd *HelloCommand) Init(provider *pipline.DependenceProvider) {

}

func (cmd *HelloCommand) Execute(request pipline.InMessage, response pipline.Response) {
	out_msg := request.MakeResponse()
	out_msg.SetText("Hello!")
	response.SendMessage(out_msg)
}

