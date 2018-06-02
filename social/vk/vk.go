package vk

import (
	vk_api "github.com/dasrmipt/go-vk-api"
	"strings"
	"github.com/ALTAIRvovan/SocNetBot/pipline"
	"github.com/FZambia/viper-lite"
	"html"
)

type VK struct {
	Client          *vk_api.VK
	commandProvider *pipline.CommandProvider
	outQueue        *chan pipline.InMessage
	adminId         int64
}

func New(config *viper.Viper) (*VK) {
	api := new(VK)
	api.Client = vk_api.New("ru")
	api.Client.Init(config.GetString("token"))
	api.adminId = config.GetInt64("admin")
	return api
}

func (api *VK) SetOutQueue(queue chan pipline.InMessage) {
	api.outQueue = &queue
}

func (api *VK) SetCommandProvider(provider *pipline.CommandProvider) {
	api.commandProvider = provider
}

func (api *VK) SetAdmin(adminId int64) {
	api.adminId = adminId
}

func (api *VK) IsAdmin(userId int64) bool {
	return api.adminId == 0 || api.adminId == userId
}

func (api *VK) Run() {
	api.Client.OnNewMessage(api.handleMessages)
	api.Client.RunLongPoll()
}

func (api *VK) handleMessages(msg *vk_api.LPMessage) {
	if msg.Flags&vk_api.FlagMessageOutBox == 0 {
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
				cmd: pipline.CmdInMessage{Cmd: com, Params: strings.Split(lines[0], " ")},
				Msg: msg,
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
