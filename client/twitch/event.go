/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package twitch

import "github.com/gempir/go-twitch-irc"

type Event struct {
	Thing
	Channel string
	User    *twitch.User
	Message *twitch.Message
}
