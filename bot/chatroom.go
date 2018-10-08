/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package bot

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/djdoeslinux/choreobot/obs_remote"
	"github.com/jinzhu/gorm"
	"strings"
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
		parts := strings.Split(result.GetResponse(), "\n")
		for _, p := range parts {
			c.client.Say(e.Channel, p)
		}
	}
}
func (c *ChatRoom) initialize(db *gorm.DB, t *client.Twitch) {
	c.client = t
	c.commands = make(map[string]command.Command)
	c.commands["!ping"] = command.Ping
	c.commands["!watchMav"] = obs_remote.Mav
	c.commands["!watchKin"] = obs_remote.Kin
	c.commands["!addCommand"] = command.AddCommand
	c.commands["!preview"] = obs_remote.TogglePreview
	c.commands["!remote"] = command.Remote
}
