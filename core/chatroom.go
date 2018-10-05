/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package core

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/jinzhu/gorm"
)

func (c *ChatRoom) handleComment(e *client.TwitchEvent) {
	//Initially this will be for questions so people don't have to keep asking them over and over.
	// should also let people bump the question from the user
	// Can also be used for a dynamic meter -- think #boom replays from DrDisrespect or KitBoga's #meme meter.
}

func (c *ChatRoom) handleCommand(e *client.TwitchEvent) {
	tokenStream := command.Tokenize(*e.Message)
	// Check if the command is already cached in our map
	key := tokenStream.GetTokenByIndex(0).String()
	cmd, _ := c.commands[key]
	if cmd == nil {
		cmd = command.NotFound
	}
	result := cmd.Evaluate(e, tokenStream)

	if result.HasResponse() {
		c.client.Say(e.Channel, result.GetResponse())
	}
}
func (b *Bot) initialize() {
	b.chats = make(map[string]*ChatRoom)
	for _, c := range b.ChatRooms {
		b.chats[c.Name] = &c
		c.initialize(b.db)
	}

}

func (c *ChatRoom) initialize(db *gorm.DB) {
	c.commands = make(map[string]command.Command)
	c.commands["!ping"] = command.Ping
	c.commands["!addCommand"] = command.AddCommand
}
