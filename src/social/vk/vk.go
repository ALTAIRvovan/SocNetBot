package vk

import (
	vk_api "github.com/dasrmipt/go-vk-api"
	"strings"
	"pipline"
)

type VK struct {
	Client          *vk_api.VK
	commandProvider *pipline.CommandProvider
	outQueue        *chan pipline.InMessage
}

func New(token string) (*VK) {
	api := new(VK)
	api.Client = vk_api.New("ru")
	api.Client.Init(token)
	return api
}

func (api *VK) SetOutQueue(queue chan pipline.InMessage) {
	api.outQueue = &queue
}

func (api *VK) SetCommandProvider(provider *pipline.CommandProvider) {
	api.commandProvider = provider
}

func (api *VK) Run() {
	api.Client.OnNewMessage(api.handleMessages)
	api.Client.RunLongPoll()
}

func (api *VK) handleMessages(msg *vk_api.LPMessage) {
	text := msg.Text
	if com, err := api.commandProvider.GetCommand(
		func(name string) bool {
			return strings.HasPrefix(text, name)
		}); err == nil {
		*api.outQueue <- &InMessage{
			cmd: com,
			msg: msg,
		}
	}
}

func (api *VK) SendOutMessage(msg pipline.OutMessage) {
	vk_msg, ok := msg.(*OutMessage)
	if !ok {
		panic("Can't send not VK message")
	}
	api.Client.Messages.Send(vk_msg.Message.ToRequestParams())
}
