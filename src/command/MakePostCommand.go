package command

import (
	"pipline"
	"github.com/dasrmipt/trello"
	vk_obj "github.com/dasrmipt/go-vk-api/obj"
	"social/vk"
	"fmt"
)

type MakePostCommand struct {
	pipline.Command
	trello_api *trello.Client
}

func (cmd *MakePostCommand) Init(provider *pipline.DependenceProvider) {
	dep := provider.GetDependence("trello")
	if trello_api, ok := dep.(*trello.Client); ok {
		cmd.trello_api = trello_api
	} else {
		panic("Dependence `trello` not found")
	}
}

func (cmd *MakePostCommand) Execute(message pipline.InMessage, response pipline.Response) {
	text := message.GetContentText()
	card := trello.Card{IDList: "5967baf92b64b3206c97d545", Name: text, Desc: text}
	err := cmd.trello_api.CreateCard(&card, trello.Defaults())
	out_msg := message.MakeResponse()
	if err == nil {
		out_msg.SetText("Card has been added " + card.ID)
		msgToAdmins := vk_obj.MessageToSend{
			PeerId: 121 + 2000000000,
			Message: fmt.Sprintf("You must make post.\nCardID: %s\n%s", card.ID, text),
		}
		response.SendMessage(&vk.OutMessage{Message: &msgToAdmins})
	} else {
		out_msg.SetText("Card has not been added")
	}
	response.SendMessage(out_msg)
}