package main

import (
	"github.com/ALTAIRvovan/SocNetBot/social/vk"
	"github.com/dasrmipt/trello"
	"github.com/FZambia/viper-lite"
	"fmt"
	"github.com/ALTAIRvovan/SocNetBot/command"
	"github.com/ALTAIRvovan/SocNetBot/pipline"
)

func initCommands(provider *pipline.CommandProvider) {
	provider.AddCommand("/hello", &command.HelloCommand{})
	provider.AddCommand("/echo_params", &command.EchoParamsCommand{})
	provider.AddCommand("/make_post", &command.MakePostCommand{})
	provider.AddCommand("/get_joke", &command.GetJokeCommand{})
	provider.AddCommand("/get_info", &command.ChatInfoCommand{})
	provider.AddCommand("#", &command.HashTagResendCommand{
		Tags: viper.GetStringMapString("tags.vk"),
	})

}

func main() {
	fmt.Print("Loading config ..... ")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	fmt.Println("OK")

	manager := pipline.NewManager()

	fmt.Print("Initing VK ..... ")
	vk_api := vk.New(viper.Sub("channels.vk"))
	manager.AddMessageChannel(vk_api)
	manager.Sender.AddSender("vk", vk_api.SendOutMessage)
	fmt.Println("OK")

	fmt.Print("Initing Trello ..... ")
	trello_api := trello.NewClient(viper.GetString("trello.key"),
		viper.GetString("trello.token"))
	fmt.Println("OK")

	fmt.Print("Initing Dependences ..... ")
	manager.DependenceProvider.AddDependence("trello", trello_api)
	manager.DependenceProvider.AddDependence("vk", vk_api)
	fmt.Println("OK")

	fmt.Print("Initing Commands ..... ")
	initCommands(manager.CommandProvider)
	fmt.Println("OK")

	manager.Run()
	fmt.Print("Bot work starts")
	vk_api.Run()
}