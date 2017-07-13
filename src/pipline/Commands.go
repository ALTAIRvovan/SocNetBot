package pipline

import "errors"

type Command interface {
	Init(producer *DependenceProvider)
	Execute(request InMessage, response Response)
}

type CommandProvider struct {
	commands     map[string]Command
	dependencies *DependenceProvider
}

func NewCommandProvider(producer *DependenceProvider) (*CommandProvider) {
	com := new(CommandProvider)
	com.commands = make(map[string]Command)
	com.dependencies = producer
	return com
}

func (com *CommandProvider) AddCommand(name string, command Command) {
	command.Init(com.dependencies)
	com.commands[name] = command
}

func (com *CommandProvider) GetCommand(comparator func(string)bool) (Command, error) {
	for name, com := range com.commands {
		if comparator(name) {
			return com, nil
		}
	}
	return nil, errors.New("Command not found")
}

// -----------------------------------------------------------------------------

type DependenceProvider struct {
	dependencies map[string]interface{}
}

func NewDependenceProvider() *DependenceProvider {
	dep := new(DependenceProvider)
	dep.dependencies = make(map[string]interface{})
	return dep
}

func (dependence *DependenceProvider) AddDependence(name string, value interface{}) {
	dependence.dependencies[name] = value
}

func (dependence *DependenceProvider) GetDependence(name string) interface{} {
	return dependence.dependencies[name]
}
