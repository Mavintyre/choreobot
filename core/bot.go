/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package core

import (
	"github.com/djdoeslinux/choreobot/client"
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
