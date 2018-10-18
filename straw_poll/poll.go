/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package straw_poll

import (
	"github.com/jinzhu/gorm"
	"time"
)

// This will implement a straw poll with a timeout or vote threshold.
// Can be used for long running polls (with names) or ad-hoc short-lived polls

var Models []interface{}

func init() {
	//Models = append(Models, &Tracker{})
}

type StrawPoll struct {
	gorm.Model
	ChannelID       uint
	Name            string
	Options         []PollOption
	RestrictedRoles []VoteRestriction
	VoteExpiresAt   time.Time
}

type VoteRestriction struct {
	gorm.Model
	StrawPollID uint
	Role        string
}

type PollOption struct {
	gorm.Model
	StrawPollID uint
}

//!newPoll <name> --option "10" --option "20" --option "30" --duration --start --voting-role
//!runPoll <name> --duration
//!modifyPoll <name> [options]
//!resetPoll <name>
