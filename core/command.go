package core

import (
	"github.com/djdoeslinux/choreobot/command"
	"github.com/gempir/go-twitch-irc"
)

func (b *bot) handleCommand(c string, u twitch.User, m twitch.Message) {
	tokenStream := command.Tokenize(m)
	// Check if the command is already cached in our map
	key := commandKey{channel: c, command: tokenStream.GetCommand()}
	cmd, _ := b.commands[key]
	if cmd == nil {
		cmd = command.NotFound
	}
	result := cmd.Evaluate(u, tokenStream)

	if result.HasResponse() {
		b.client.Say(c, result.GetResponse())
	}
}
