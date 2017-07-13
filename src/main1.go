package main

import (
	"github.com/dasrmipt/go-vk-api"
	vk_obj "github.com/dasrmipt/go-vk-api/obj"
	"log"
	"strconv"
	"strings"
	"github.com/dasrmipt/trello"
	cmd "./command"
	"fmt"
)

type VKCmdMessage struct {
	cmd.Message
	api *vk.VK
	msg *vk.LPMessage
}

func (msg VKCmdMessage) SendResponse(response string) {
	msg.api.Messages.Send(vk.RequestParams{
		"peer_id":          strconv.FormatInt(msg.msg.FromID, 10),
		"message":          response,
		"forward_messages": strconv.FormatInt(msg.msg.ID, 10),
	})
}

func (msg VKCmdMessage) GetText() string {
	return msg.msg.Text
}

func handleVKMessage(channel chan cmd.Message, api *vk.VK, msg *vk.LPMessage) {
	if msg.Flags&vk.FlagMessageOutBox == 0 {
		if !strings.HasPrefix(msg.Text, "/") {
			return
		}
		cmd_message := VKCmdMessage{api: api, msg: msg}
		text := msg.Text
		if strings.HasPrefix(text, "/hello") {
			cmd_message.SendResponse("HELLO!")
		}
		channel <- cmd_message
	}
}

func handleQueueMessages(channel chan cmd.Message, vkQueue chan vk_obj.MessageToSend) {
	trello_api := trello.NewClient("bb",
		"vdd")
	for msg := range channel {
		msgText := msg.GetText()
		if strings.HasPrefix(msgText, "/get_boards") {
			organization, error := trello_api.GetOrganization("dasrmipt", trello.Defaults())
			if error != nil {
				fmt.Print("Organization not found")
				return
			}
			boards, error := organization.GetBoards(trello.Defaults())
			if error != nil {
				fmt.Print("Board not found")
				continue
			}
			response := "Board found: \n"
			for _, board := range boards {
				response += fmt.Sprintf("%s \n", board.Name)
			}
			msg.SendResponse(response)
		}

		if strings.HasPrefix(msgText, "/board_lists") {
			board, error := trello_api.GetBoard("mr6QUygd", trello.Defaults())
			if error != nil {
				fmt.Print("Board not found")
				continue
			}
			lists, error := board.GetLists(trello.Defaults())
			if error != nil {
				fmt.Print("Lists not found")
				continue
			}
			response := "Board found: \n"
			for _, list := range lists {
				response += fmt.Sprintf("%s %s\n", list.Name, list.ID)
			}
			msg.SendResponse(response)
		}

		if strings.HasPrefix(msgText, "/make_post") {
			text := strings.TrimPrefix(msgText, "/make_post")
			card := trello.Card{IDList: "5967baf92b64b3206c97d545", Name: text, Desc: text}
			error := trello_api.CreateCard(&card, trello.Defaults())
			if error == nil {
				msg.SendResponse("Card has been added " + card.ID)
				vkQueue <- vk_obj.MessageToSend{
					PeerId: 121 + 2000000000,
					Message: fmt.Sprintf("You must make post.\nCardID: %s\n%s", card.ID, text),
				}
			} else {
				msg.SendResponse("Card has not been added")
			}

		}

		if strings.HasPrefix(msgText, "/card_info") {
			text := strings.TrimPrefix(msgText, "/card_info ")
			words := strings.Fields(text)
			card, error := trello_api.GetCard(words[0], trello.Arguments{"actions": "all"})
			if error == nil {
				msg.SendResponse(fmt.Sprintf("Card has been found \nName:%s\nID:%s\nStatus:%b", card.Name, card.ID, card.Closed))
			} else {
				msg.SendResponse("Card has not been added")
			}
		}
	}
}

func sendVKMessages(channel chan vk_obj.MessageToSend, api *vk.VK) {
	for msg := range channel {
		api.Messages.Send(msg.ToRequestParams())
	}
}

func main() {
	api := vk.New("ru")
	// set http proxy
	//api.Proxy = "localhost:8080"

	vk2trelloQueue := make(chan cmd.Message)
	trello2vkQueue := make(chan vk_obj.MessageToSend)
	go handleQueueMessages(vk2trelloQueue, trello2vkQueue)
	go sendVKMessages(trello2vkQueue, api)

	err := api.Init("bbb")

	if err != nil {
		log.Fatalln(err)
	}

	api.OnNewMessage(func(msg *vk.LPMessage) { handleVKMessage(vk2trelloQueue, api, msg) })

	api.RunLongPoll()
}
