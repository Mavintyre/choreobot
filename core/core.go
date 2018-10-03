package core

import (
	"github.com/djdoeslinux/choreobot/command"
	"github.com/djdoeslinux/choreobot/moderator"
	"github.com/gempir/go-twitch-irc"
)

type bot struct {
	username   string
	commands   map[commandKey]command.Command
	client     twitch.Client
	moderators map[string]moderator.Moderator
}

type commandKey struct {
	channel string
	command string
}

func (b *bot) Start() {

}

// Implement the callbacks for the twitch irc library
func (b *bot) GetUserName() string {
	return b.username
}

func (b *bot) GetToken() string {
	panic("implement me")
}

func (b *bot) GetChannelList() []string {
	panic("implement me")
}

func (b *bot) OnConnect() {
	// This function intentionally left blank
}

func (b *bot) OnWhisper(u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *bot) OnMessage(c string, u twitch.User, m twitch.Message) {
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

func (b *bot) OnRoomState(c string, u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *bot) OnClearChat(c string, u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *bot) OnUserNotice(c string, u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *bot) OnUserState(c string, u twitch.User, m twitch.Message) {
	panic("implement me")
}

func (b *bot) moderateMessage(c string, user twitch.User, message twitch.Message) {
	mod, _ := b.moderators[c]
	if mod == nil {
		return
	}
	//TODO: Decide if we should block on this or not.
	mod.Moderate(user, message)
}

func (b *bot) handleComment(c string, user twitch.User, message twitch.Message) {
	//Initially this will be for questions so people don't have to keep asking them over and over.
	// should also let people bump the question from the user
	// Can also be used for a dynamic meter -- think #boom replays from DrDisrespect or KitBoga's #meme meter.
}

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
