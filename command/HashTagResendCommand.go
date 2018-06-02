package command

import "github.com/ALTAIRvovan/SocNetBot/pipline"
import (
	vk_obj "github.com/dasrmipt/go-vk-api/obj"
	"strconv"
	"fmt"
	vk "github.com/ALTAIRvovan/SocNetBot/social/vk"
)

type HashTagResendCommand struct {
	pipline.Command
	Tags map[string]string
	vk_api *vk.VK
}

func (cmd *HashTagResendCommand) Init(provider *pipline.DependenceProvider) {
	dep := provider.GetDependence("vk")
	if vk_api, ok := dep.(*vk.VK); ok {
		cmd.vk_api = vk_api
	} else {
		panic("Dependence `vk` not found")
	}
}

func (cmd *HashTagResendCommand) Execute(request pipline.InMessage, response pipline.Response) {
	if vkMsg, ok := request.(*vk.InMessage); ok {

		params := request.GetCmd().Params
		peerIdStr, ok := cmd.Tags[params[0]]
		if !ok {
			return
		}
		peerId, err := strconv.Atoi(peerIdStr)
		if err != nil {
			return
		}
		msg := vk_obj.MessageToSend{
			PeerId: int64(peerId),
		}

		authorId := vkMsg.Msg.GetAuthor()
		fromId := vkMsg.Msg.FromID

		if fromId == int64(peerId) {
			return
		}

		if authorId != fromId {
			params := map[string]string {
				"chat_id": strconv.FormatInt(fromId - 2000000000, 10),
			}
			chat, err := cmd.vk_api.Client.Messages.GetChat(params)
			if err != nil {
				return
			}
			msg.Message = fmt.Sprintf(fmt.Sprintf("Сообщение из чата: %s", chat.Title))
		}
		msg.FwdMessages = []int64{vkMsg.Msg.ID}
		response.SendMessage(&vk.OutMessage{Message: &msg})

		if authorId == fromId {
			out_msg := request.MakeResponse()
			out_msg.SetText("Ваше сообщение переслано")
			response.SendMessage(out_msg)
		}
	}
}
