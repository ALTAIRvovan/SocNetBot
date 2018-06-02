package command

import "github.com/ALTAIRvovan/SocNetBot/pipline"
import (
	vk "github.com/ALTAIRvovan/SocNetBot/social/vk"
	"fmt"
);

type ChatInfoCommand struct {
	pipline.Command
	vk_api *vk.VK
}

func (cmd *ChatInfoCommand) Init(provider *pipline.DependenceProvider) {
	dep := provider.GetDependence("vk")
	if vk_api, ok := dep.(*vk.VK); ok {
		cmd.vk_api = vk_api
	} else {
		panic("Dependence `vk` not found")
	}
}

func (cmd *ChatInfoCommand) Execute(request pipline.InMessage, response pipline.Response) {
	out_msg := request.MakeResponse()
	if vk_request, ok := request.(*vk.InMessage); ok {
		author_id := vk_request.Msg.GetAuthor()
		if cmd.vk_api.IsAdmin(author_id) {
			//fmt.Print(vk_request.Msg)
			out_msg.SetText(fmt.Sprintf("Message sended from chat: %d", vk_request.Msg.FromID))
		} else {
			out_msg.SetText("It can be only by admin")
		}
	} else {
		out_msg.SetText("It work only in VK!")
	}
	response.SendMessage(out_msg)
}

