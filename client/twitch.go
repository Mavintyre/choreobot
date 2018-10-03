package client

import (
	"fmt"
	"github.com/gempir/go-twitch-irc"
	"strings"
)

type BotConfig interface {
	GetUserName() string
	GetToken() string
	GetChannelList() []string
	OnConnect()
	OnWhisper(twitch.User, twitch.Message)
	OnMessage(string, twitch.User, twitch.Message)
	OnRoomState(string, twitch.User, twitch.Message)
	OnClearChat(string, twitch.User, twitch.Message)
	OnUserNotice(string, twitch.User, twitch.Message)
	OnUserState(string, twitch.User, twitch.Message)
}

func connect(b BotConfig) *twitch.Client {
	token := b.GetToken()
	if !strings.HasPrefix(token, "oauth:") {
		token = fmt.Sprintf("oauth:%s", token)
	}
	client := twitch.NewClient(b.GetUserName(), token)
	client.OnConnect(b.OnConnect)
	client.OnNewWhisper(b.OnWhisper)
	client.OnNewMessage(b.OnMessage)
	client.OnNewRoomstateMessage(b.OnRoomState)
	client.OnNewClearchatMessage(b.OnClearChat)
	client.OnNewUsernoticeMessage(b.OnUserNotice)
	client.OnNewUserstateMessage(b.OnUserState)
	for _, channelName := range b.GetChannelList() {
		client.Join(channelName)
	}
	go func() {
		err := client.Connect()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return client
}
