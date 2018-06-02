package vk

import (
	"github.com/ALTAIRvovan/SocNetBot/pipline"
	vk_api "github.com/dasrmipt/go-vk-api"
	vk_obj "github.com/dasrmipt/go-vk-api/obj"
)

type InMessage struct {
	pipline.InMessage
	cmd pipline.CmdInMessage
	Msg *vk_api.LPMessage
}

type OutMessage struct {
	pipline.OutMessage
	Message *vk_obj.MessageToSend
}

func (request *InMessage) GetCmd() pipline.CmdInMessage {
	return request.cmd
}

func (msg *OutMessage) GetType() string {
	return "vk"
}

func (msg *InMessage) MakeResponse() pipline.OutMessage {
	message := new(vk_obj.MessageToSend)
	message.PeerId = msg.Msg.FromID
	message.FwdMessages = []int64{msg.Msg.ID}
	return &OutMessage{Message: message}
}

func (msg *InMessage) GetContentText() string {
	return msg.Msg.Text
}

func (msg *OutMessage) SetText(text string) {
	msg.Message.Message = text
}