/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package command

import "github.com/gempir/go-twitch-irc"

var NotFound Command

//AddCommand - create a command
//Permit -
//Grant -
//DropCommand -
//Snapshot - Save the current config/state as a restorable thing
//Restore -
//Help -
//Mute - Supress all responses from the bot, but continue moderating

func init() {
	NotFound = Command(notFound{})
}

type notFound struct{}

func (notFound) IsAllowed(u twitch.User) bool {
	// Yes everyone can get an error message
	return true
}

func (notFound) Evaluate(u twitch.User, t TokenStream) Result {
	panic("implement me")
	//Return default not found message
}
