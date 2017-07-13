package command

import (
	"github.com/dnaeon/go-chucknorris/api"
	"pipline"
)

type GetJokeCommand struct {
	pipline.Command
	joke_api *api.Client
}

func (cmd *GetJokeCommand) Init(provider *pipline.DependenceProvider) {
	cmd.joke_api = api.NewClient(nil)
}

func (cmd *GetJokeCommand) Execute(message pipline.InMessage, response pipline.Response) {
	joke, err := cmd.joke_api.RandomJoke()
	out_msg := message.MakeResponse()
	if err == nil {
		out_msg.SetText(joke.Value)
	} else {
		out_msg.SetText("Joke not found :(")
	}
	response.SendMessage(out_msg)
}