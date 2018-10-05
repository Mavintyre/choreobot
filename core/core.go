/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package core

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/djdoeslinux/choreobot/moderator"
	"github.com/gempir/go-twitch-irc"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sanity-io/litter"
)

var Models []interface{}

func init() {
	Models = append(Models, &Bot{}, &ChatRoom{}, &MessageHandler{})
}

type Bot struct {
	gorm.Model
	UserName   string
	OAuthToken string
	ChatRooms  []ChatRoom

	// Private members down here
	client     *client.Twitch
	db         *gorm.DB
	moderators map[string]moderator.Moderator
	commands   map[commandKey]command.Command
	chats      map[string]*ChatRoom
}

type ChatRoom struct {
	gorm.Model
	BotID           uint
	IsEnabled       bool
	Name            string
	MessageHandlers []MessageHandler

	//Private members down here
	isModerator bool
}

type MessageHandler struct {
	gorm.Model
	ChannelID  uint
	Namespace  string
	Name       string
	IsDisabled bool
}

type commandKey struct {
	channel string
	command string
}

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
	c := e.Channel
	u := e.User
	m := e.Message
	//We always want to moderate all message regardless
	b.moderateMessage(c, *u, *m)

	//Handle the message differently based on the first character
	textAsBytes := []byte(m.Text)
	switch textAsBytes[0] {
	case '!':
		b.handleCommand(e)
	case '#':
		b.handleComment(c, *u, *m)
	}
}

func (b *Bot) moderateMessage(c string, user twitch.User, message twitch.Message) {
	mod, exists := b.moderators[c]
	if !exists {
		return
	}
	//TODO: Decide if we should block on this or not.
	mod.Moderate(user, message)
}

func (b *Bot) handleComment(c string, user twitch.User, message twitch.Message) {
	//Initially this will be for questions so people don't have to keep asking them over and over.
	// should also let people bump the question from the user
	// Can also be used for a dynamic meter -- think #boom replays from DrDisrespect or KitBoga's #meme meter.
}

func (b *Bot) handleCommand(e *client.TwitchEvent) {
	tokenStream := command.Tokenize(*e.Message)
	// Check if the command is already cached in our map
	key := commandKey{channel: e.Channel, command: tokenStream.GetTokenByIndex(0).String()}
	cmd, _ := b.commands[key]
	if cmd == nil {
		cmd = command.NotFound
	}
	result := cmd.Evaluate(e, tokenStream)

	if result.HasResponse() {
		b.client.Say(e.Channel, result.GetResponse())
	}
}
func (b *Bot) initialize() {
	b.commands = make(map[commandKey]command.Command)
	b.moderators = make(map[string]moderator.Moderator)
	b.chats = make(map[string]*ChatRoom)
	for _, c := range b.ChatRooms {
		b.chats[c.Name] = &c
		b.commands[commandKey{c.Name, "!ping"}] = command.GetPing()
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
