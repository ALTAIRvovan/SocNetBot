package vk

import "github.com/ALTAIRvovan/SocNetBot/pipline"

type Request struct{
	pipline.Request
	cmd *pipline.Command
	msg *pipline.InMessage
}

func (request *Request) GetCmd() *pipline.Command {
	return request.cmd
}


func (request *Request) GetMessage() *pipline.InMessage {
	return request.msg
}
