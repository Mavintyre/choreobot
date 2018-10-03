package core

import (
	"github.com/djdoeslinux/choreobot/command"
	"github.com/gempir/go-twitch-irc"
)

type bot struct {
	username string
	commands map[commandKey]command.Command
	client   twitch.Client
}

type commandKey struct {
	channel string
	command string
}

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

func (b *bot) moderateMessage(s string, user twitch.User, message twitch.Message) {

}
