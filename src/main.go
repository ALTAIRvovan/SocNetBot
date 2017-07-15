package main

import "pipline"
import (
	"social/vk"
	"command"
	"github.com/dasrmipt/trello"
	"github.com/FZambia/viper-lite"
)

func initCommands(provider *pipline.CommandProvider) {
	provider.AddCommand("/hello", &command.HelloCommand{})
	provider.AddCommand("/echo_params", &command.EchoParamsCommand{})
	provider.AddCommand("/make_post", &command.MakePostCommand{})
	provider.AddCommand("/get_joke", &command.GetJokeCommand{})
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	manager := pipline.NewManager()

	vk_api := vk.New(viper.GetString("channels.vk.token"))
	manager.AddMessageChannel(vk_api)
	manager.Sender.AddSender("vk", vk_api.SendOutMessage)

	trello_api := trello.NewClient(viper.GetString("trello.key"),
		viper.GetString("trello.token"))


	manager.DependenceProvider.AddDependence("trello", trello_api)
	manager.DependenceProvider.AddDependence("vk", vk_api.Client)
	initCommands(manager.CommandProvider)

	manager.Run()
	vk_api.Run()

}