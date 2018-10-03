package turing_test

import (
	"github.com/djdoeslinux/choreobot/command"
	"github.com/gempir/go-twitch-irc"
)

//This is a generic responder command implementation. It is configured with

type responder struct {
	templates []string
}

type Responder interface {
	command.Command
	GetResponseByIndex(i int) string
	GetRandomResponse() string
	AddResponse(string) int
}

func (*responder) Evaluate(u twitch.User, t command.TokenStream) command.Result {
	panic("implement me")

}
