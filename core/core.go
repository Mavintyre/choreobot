/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package core

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sanity-io/litter"
)

func (b *Bot) Start(db *gorm.DB) {
	b.db = db
	b.initialize()
	b.client = client.NewTwitchClient(b.UserName)
	var chats []string
	for c, _ := range b.chats {
		chats = append(chats, c)
	}
	b.client.Start(b, chats...)
	eventChan := b.client.GetEventChannel()

	for event := range eventChan {
		litter.Dump(event)
		switch event.Thing {
		case client.Message:
			b.handleMessage(event)
		case client.Whisper:

		}
	}
}

func (b *Bot) Stop() {

}

func LoadBot(db gorm.DB, name string) (b *Bot, err error) {
	b = &Bot{}
	db.FirstOrInit(b, Bot{UserName: name})

	return
}

// Implement client.Authenticator interface
func (b *Bot) GetToken() string {
	return b.OAuthToken
}

func (b *Bot) handleMessage(e *client.TwitchEvent) {
	c := b.chats[e.Channel]
	m := e.Message
	//We always want to moderate all message regardless
	c.Moderator.Moderate(e)

	//Handle the message differently based on the first character
	textAsBytes := []byte(m.Text)
	switch textAsBytes[0] {
	case '!':
		c.handleCommand(e)
	case '#':
		c.handleComment(e)
	}
}

func (b *Bot) JoinNewChat(c string) {
	if _, exists := b.chats[c]; exists {
		return
	}
	//TODO: Setup the default parameters for the room.
	newChat := &ChatRoom{Name: c, IsEnabled: true, BotID: b.ID}
	b.db.Create(newChat)
	b.client.Join(c)

}

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
