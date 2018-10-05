/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package core

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/djdoeslinux/choreobot/moderator"
	"github.com/jinzhu/gorm"
)

var Models []interface{}

func init() {
	Models = append(Models, &Bot{}, &ChatRoom{}, &MessageHandler{}, &Permission{}, &Role{})
}

type Bot struct {
	gorm.Model
	UserName   string
	OAuthToken string
	ChatRooms  []ChatRoom

	// Private members down here
	client *client.Twitch
	db     *gorm.DB
	chats  map[string]*ChatRoom
}

type ChatRoom struct {
	gorm.Model
	BotID           uint
	IsEnabled       bool
	Name            string
	Moderator       *moderator.Moderator
	MessageHandlers []MessageHandler
	Permissions     []Permission

	//Private members down here
	isModerator bool
	commands    map[string]command.Command
	client      *client.Twitch
}

type MessageHandler struct {
	gorm.Model
	ChannelID               uint
	Namespace               string
	Name                    string
	CommandImplementationID uint
	IsDisabled              bool
}

type Permission struct {
	gorm.Model
	ChannelID uint
	RoleID    uint
	CommandID uint
	Priority  int
	Grant     string
}

type Role struct {
	gorm.Model
	ChannelID uint
	Name      string
}
