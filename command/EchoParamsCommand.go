package command

import (
	"github.com/ALTAIRvovan/SocNetBot/pipline"
	"fmt"
)

type EchoParamsCommand struct {
	pipline.Command
}

func (cmd *EchoParamsCommand) Init(provider *pipline.DependenceProvider) {

}

func (cmd *EchoParamsCommand) Execute(message pipline.InMessage, response pipline.Response) {
	out_msg := message.MakeResponse()
	out_text := "Params:\n"
	for index, param := range message.GetCmd().Params {
		out_text += fmt.Sprintf("%d: %s\n", index, param)
	}
	out_msg.SetText(out_text)
	response.SendMessage(out_msg)
}
