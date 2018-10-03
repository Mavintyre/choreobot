package command

import (
	"github.com/gempir/go-twitch-irc"
)

type Command interface {
	IsAllowed(u twitch.User) bool
	Evaluate(u twitch.User, t TokenStream) Result
}

type Result interface {
	HasResponse() bool
	GetResponse() string
}

type Token string

type TokenStream interface {
	GetCommand() string
	GetRaw() twitch.Message
	NumArgs() int
	GetToken(index int) Token
	NotDone() bool
	Next()
	CurrentToken() Token
}

func Tokenize(m twitch.Message) TokenStream {
	//TODO: Choose a proper lexxer but for now just split on whitespace.
	return newStupidLexxer(m)
}
