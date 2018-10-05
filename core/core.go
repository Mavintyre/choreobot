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
)

type Bot struct {
	gorm.Model
	UserName   string
	OAuthToken string
	ChatRooms  []ChatRoom

	// Private members down here
	client     *client.Twitch
	dbPath     string
	moderators map[string]moderator.Moderator
	commands   map[commandKey]command.Command
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
	command command.Token
}

func (b *Bot) Start() {
	b.client = client.NewTwitchClient(b.UserName)
	b.client.Start(b)
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

func (b *Bot) OnConnect() {
	// This function intentionally left blank
}

func (b *Bot) OnWhisper(u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *Bot) OnMessage(c string, u twitch.User, m twitch.Message) {
	//We always want to moderate all message regardless
	b.moderateMessage(c, u, m)

	//Handle the message differently based on the first character
	textAsBytes := []byte(m.Text)
	switch textAsBytes[0] {
	case '!':
		b.handleCommand(c, u, m)
	case '#':
		b.handleComment(c, u, m)
	}
}

func (b *Bot) OnRoomState(c string, u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *Bot) OnClearChat(c string, u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *Bot) OnUserNotice(c string, u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *Bot) OnUserState(c string, u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *Bot) moderateMessage(c string, user twitch.User, message twitch.Message) {
	mod, _ := b.moderators[c]
	if mod == nil {
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

func (b *Bot) handleCommand(c string, u twitch.User, m twitch.Message) {
	tokenStream := command.Tokenize(m)
	// Check if the command is already cached in our map
	key := commandKey{channel: c, command: tokenStream.GetTokenByIndex(0)}
	cmd, _ := b.commands[key]
	if cmd == nil {
		cmd = command.NotFound
	}
	result := cmd.Evaluate(u, tokenStream)

	if result.HasResponse() {
		b.client.Say(c, result.GetResponse())
	}
}
