/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package autoregister

import (
	"github.com/djdoeslinux/choreobot/bot"
	"github.com/djdoeslinux/choreobot/command/counter"
	"github.com/djdoeslinux/choreobot/command/loyalty_points"
	"github.com/djdoeslinux/choreobot/command/turing_test"
	"github.com/djdoeslinux/choreobot/meter"
	"github.com/djdoeslinux/choreobot/moderator"
	"github.com/djdoeslinux/choreobot/straw_poll"
	"github.com/djdoeslinux/choreobot/user"
)

func init() {
	bot.Register()
	counter.Register()
	loyalty_points.Register()
	meter.Register()
	moderator.Register()
	straw_poll.Register()
	turing_test.Register()
	user.Register()
}
