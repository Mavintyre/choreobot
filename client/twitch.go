/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package client

import (
	"fmt"
	"github.com/gempir/go-twitch-irc"
	"strings"
)

type Twitch struct {
	user string
	*twitch.Client
	eventChan  chan *TwitchEvent
	actionChan chan action
}

type Authenticator interface {
	GetToken() string
}

type Thing int

const (
	_ Thing = iota
	Connect
	Whisper
	Message
	RoomState
	ChatClear
	UserNotice
	UserState
)

type TwitchEvent struct {
	Thing
	Channel string
	User    twitch.User
	Message twitch.Message
}

type action struct {
}

func NewTwitchClient(username string) *Twitch {
	t := &Twitch{user: username}
	t.eventChan = make(chan *TwitchEvent)
	t.actionChan = make(chan action)
	return t
}

func (t *Twitch) GetEventChannel() chan *TwitchEvent {
	return t.eventChan
}

func (t *Twitch) Start(a Authenticator, channels ...string) {
	// Connect
	tok := a.GetToken()
	if !strings.HasPrefix(tok, "oauth:") {
		tok = fmt.Sprintf("oauth:%s", tok)
	}
	t.Client = twitch.NewClient(t.user, tok)

	t.Client.OnConnect(func() { handleConnect(t) })
	t.Client.OnNewWhisper(func(u twitch.User, m twitch.Message) { handleWhisper(t, u, m) })
	t.Client.OnNewMessage(func(c string, u twitch.User, m twitch.Message) { handleMessage(t, c, u, m) })
	t.Client.OnNewRoomstateMessage(func(c string, u twitch.User, m twitch.Message) { handleRoomState(t, c, u, m) })
	t.Client.OnNewClearchatMessage(func(c string, u twitch.User, m twitch.Message) { handleChatClear(t, c, u, m) })
	t.Client.OnNewUsernoticeMessage(func(c string, u twitch.User, m twitch.Message) { handleUserNotice(t, c, u, m) })
	t.Client.OnNewUserstateMessage(func(c string, u twitch.User, m twitch.Message) { handleUserState(t, c, u, m) })

	for _, channelName := range channels {
		t.Client.Join(channelName)
	}
	err := t.Client.Connect()
	if err != nil {
		fmt.Println(err)
	}
	t.Stop()

}

func (t *Twitch) Stop() {

}

func handleConnect(t *Twitch) {
	t.eventChan <- &TwitchEvent{Connect, "", nil, nil}
}
func handleWhisper(t *Twitch, user twitch.User, message twitch.Message) {
	t.eventChan <- &TwitchEvent{Whisper, "", user, message}

}
func handleMessage(t *Twitch, s string, user twitch.User, message twitch.Message) {
	t.eventChan <- &TwitchEvent{Message, s, user, message}

}
func handleRoomState(t *Twitch, s string, user twitch.User, message twitch.Message) {
	t.eventChan <- &TwitchEvent{RoomState, s, user, message}

}
func handleChatClear(t *Twitch, s string, user twitch.User, message twitch.Message) {
	t.eventChan <- &TwitchEvent{ChatClear, s, user, message}

}
func handleUserNotice(t *Twitch, s string, user twitch.User, message twitch.Message) {
	t.eventChan <- &TwitchEvent{UserNotice, s, user, message}

}
func handleUserState(t *Twitch, s string, user twitch.User, message twitch.Message) {
	t.eventChan <- &TwitchEvent{UserState, s, user, message}
}
