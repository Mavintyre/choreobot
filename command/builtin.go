package command

import "github.com/gempir/go-twitch-irc"

var NotFound Command

func init() {
	NotFound = Command(notFound{})
}

type notFound struct{}

func (notFound) Evaluate(u twitch.User, t TokenStream) Result {
	panic("implement me")
	//Return default not found message
}
