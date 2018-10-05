/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package command

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/gempir/go-twitch-irc"
)

var NotFound Command

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
	NotFound = Command(notFound{})
}

type notFound struct{}

func (notFound) IsAllowed(u twitch.User) bool {
	// Yes everyone can get an error message
	return true
}

func (notFound) Evaluate(e *client.TwitchEvent, t TokenStream) Result {
	panic("implement me")
	//Return default not found message
}

type builtin struct {
	cmd string
	exe func(e *client.TwitchEvent, s TokenStream) Result
}

func GetPing() Command {
	return &builtin{"ping", ping}
}

func (b *builtin) Evaluate(e *client.TwitchEvent, s TokenStream) Result {
	return b.exe(e, s)
}

func (b *builtin) IsAllowed(u twitch.User) bool {
	return true
}

func ping(event *client.TwitchEvent, stream TokenStream) Result {
	return &Reply{"PONG!"}
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
