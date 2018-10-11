/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package builtins

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/djdoeslinux/choreobot/command/counter"
	"github.com/djdoeslinux/choreobot/command/scoreboard"
	"github.com/djdoeslinux/choreobot/command/turing_test"
	"github.com/gempir/go-twitch-irc"
)

var NotFound command.Command

//AddCommand - create a command
//Permit -
//Grant -
//DropCommand -
//Snapshot - Save the current config/state as a restorable thing
//Restore -
//Help -
//Mute - Supress all responses from the bot, but continue moderating
//ping - respond with pong

func init() {
	NotFound = command.Command(notFound{})
}

type notFound struct{}

func (notFound) IsAllowed(u twitch.User) bool {
	// Yes everyone can get an error message
	return true
}

func (notFound) Evaluate(e *client.TwitchEvent, t command.TokenStream) command.Result {
	panic("implement me")
	//Return default not found message
}

var AddCommand command.Command
var Permit command.Command
var Grant command.Command
var DropCommand command.Command
var Snapshot command.Command
var Restore command.Command
var Help command.Command
var Mute command.Command
var Ping command.Command

func init() {
	AddCommand = &builtin{"AddCommand", addCommand}
	//Permit = &builtin{"Permit", permit}
	//Grant = &builtin{"Grant", grant}
	//DropCommand = &builtin{"DropCommand", dropCommand}
	//Snapshot = &builtin{"Snapshot", snapshot}
	//Restore = &builtin{"Restore", restore}
	//Help = &builtin{"Help", help}
	//Mute = &builtin{"Mute", mute}
	Ping = &builtin{"ping", ping}
}

type builtin struct {
	cmd string
	exe func(e *client.TwitchEvent, s command.TokenStream) command.Result
}

func (b *builtin) Evaluate(e *client.TwitchEvent, s command.TokenStream) command.Result {
	return b.exe(e, s)
}

func (b *builtin) IsAllowed(u twitch.User) bool {
	return true
}

func ping(event *client.TwitchEvent, stream command.TokenStream) command.Result {
	return &Reply{"PONG!"}
}

func addCommand(e *client.TwitchEvent, s command.TokenStream) command.Result {
	s.Seek(1)
	name, err := s.PopToken()
	if err != nil {

		return TODO //Return a usage message
	}
	namespace, err := s.PopToken()
	if err != nil {
		return TODO	//Return a usage message
	}

	switch namespace.String() {
	case "quote", "response":
		t := turing_test.NewBlankTuring()
		t.Name = name.String()
		return TODO //Parse and respond
	case "counter":
		counter.NewBlankCounter()
		return TODO //Parse and respond
	case "scoreboard":
		scoreboard.NewScoreboard()
	default:
		return TODO //Return a no implementation error
	}

	return &Reply{"no command for you"}
}

type Reply struct {
	Value string
}

func (r *Reply) HasResponse() bool {
	if r.Value != "" {
		return true
	}
	return false
}

func (r *Reply) GetResponse() string {
	return r.Value
}
