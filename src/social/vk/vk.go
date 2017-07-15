package vk

import (
	vk_api "github.com/dasrmipt/go-vk-api"
	"strings"
	"pipline"
	"github.com/urShadow/go-vk-api"
	"html"
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
	if msg.Flags&vk.FlagMessageOutBox == 0 {
		msg.Text = html.UnescapeString(msg.Text)
		msg.Text = strings.Replace(msg.Text, "<br>", "\n", -1)
		lines := strings.SplitN(msg.Text, "\n", 2)
		if len(lines) > 1 {
			msg.Text = lines[1]
		} else {
			msg.Text = ""
		}
		if com, err := api.commandProvider.GetCommand(
			func(name string) bool {
				return strings.HasPrefix(lines[0], name)
			}); err == nil {
			*api.outQueue <- &InMessage{
				cmd: pipline.CmdInMessage{Cmd:com, Params: strings.Split(lines[0], " ")},
				msg: msg,
			}
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
