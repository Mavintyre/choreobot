package turing_test

import (
	"github.com/djdoeslinux/choreobot/command"
	"github.com/gempir/go-twitch-irc"
)

//This is a generic responder command implementation.
// It can handle any rote stateless responses by template.
// Use cases:
// Simple Response: !rules -- Just says something verbatim in the chat (simple response is actually just a 1 template multiresponse)
// MultiResponse: !quote [int|tag] -- Keeps track of multiple possible responses. Picks randomly unless a specific index is requested
///// MultiResponse
// Templated: !law {U} {1} {2} {...} -- Interpolates a response based on a template. Arguments are space separated. An ellipsis indicates "all remaining arguments"

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
