package vk

import (
	"pipline"
	vk_api "github.com/dasrmipt/go-vk-api"
	vk_obj "github.com/dasrmipt/go-vk-api/obj"
)

type BaseMessage struct {

}

type InMessage struct {
	BaseMessage
	pipline.InMessage
	cmd pipline.Command
	msg *vk_api.LPMessage
}

type OutMessage struct {
	pipline.OutMessage
	BaseMessage
	Message *vk_obj.MessageToSend
}

func (request *InMessage) GetCmd() pipline.Command {
	return request.cmd
}

func (msg *OutMessage) GetType() string {
	return "vk"
}

func (msg *InMessage) MakeResponse() pipline.OutMessage {
	message := new(vk_obj.MessageToSend)
	message.PeerId = msg.msg.FromID
	message.FwdMessages = []int64{msg.msg.ID}
	return &OutMessage{Message: message}
}

func (msg *InMessage) GetContentText() string {
	return msg.msg.Text
}

func (msg *OutMessage) SetText(text string) {
	msg.Message.Message = text
}