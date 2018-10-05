/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package moderator

import (
	"github.com/djdoeslinux/choreobot/client"
	"github.com/gempir/go-twitch-irc"
	"github.com/jinzhu/gorm"
)

var Models []interface{}

func init() {
	Models = append(Models, &Moderator{}, &Rule{})
}

type ModeratorI interface {
	Moderate(u twitch.User, m twitch.Message)
}

type ConfigurableModerator interface {
	ModeratorI
	AddRule(r Rule)
	DisableRule(r Rule)
	ModifyRule(r Rule)
	GetRules() []Rule
	GetRuleByLabel(n string) Rule
}

type Rule struct {
	gorm.Model
	Name        string
	ModeratorID uint
}

type Moderator struct {
	gorm.Model
	ChannelID uint
	Rules     []Rule
}

type ModerationResult struct {
}

func (m *Moderator) Moderate(e *client.TwitchEvent) ModerationResult {
	return ModerationResult{}
}

// Features to be implemented
//   Blacklisting
//	 Whitelist
//	 Permit
//	 Profanity avoidance
//// Quotas and Limits
//	 Caps Limit
//	 Repetition Spam
//	 Emote Limit
//   Symbols (?) Limit

// Limits should have decaying quotas and should probably take some sort of karma score into account.
