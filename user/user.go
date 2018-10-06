/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package user

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/jinzhu/gorm"
	"time"
)

// This will track a user across channels
var Models []interface{}

func init() {
	Models = append(Models, &User{}, &UserChatState{})
}

type User struct {
	gorm.Model
	Name           string
	TwitchID       int
	UserChatStates []UserChatState
}

type UserChatState struct {
	gorm.Model
	UserID    uint
	ChannelID uint
	LastSeen  time.Time
	BanCount  int
	LastBan   int
}

func GetUserByEvent(e *client.TwitchEvent) *User {
	panic("implement me")
}

func GetUserByName(n string) *User {
	return nil
}
