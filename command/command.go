/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package command

import (
	"fmt"
	"github.com/djdoeslinux/choreobot/client"
	"github.com/gempir/go-twitch-irc"
)

var TODO Result

func init() {
	TODO = &tt{}
}

type Command interface {
	IsAllowed(u twitch.User) bool
	Evaluate(e *client.TwitchEvent, t TokenStream) Result
}

type Result interface {
	HasResponse() bool
	GetResponse() string
}

type Token interface {
	fmt.Stringer
}

type TokenStream interface {
	GetRaw() twitch.Message
	NumTokens() int
	GetTokenByIndex(index int) Token
	NotDone() bool
	Next() error
	Prev() error
	Token() Token
	PopToken() (Token, error) //
	Seek(int) error
}

type CommandStream interface {
	TokenStream
	GetCommand() string
}

func Tokenize(m twitch.Message) TokenStream {
	//TODO: Choose a proper lexxer but for now just split on whitespace.
	return newStupidLexxer(m)
}

func Error(booBoos ...interface{}) Result {
	panic("x")
}

type tt struct {
}

func (*tt) HasResponse() bool {
	return false
}

func (*tt) GetResponse() string {
	return "Ok"
}
