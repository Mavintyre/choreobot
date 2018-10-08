/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package registry

import (
	"github.com/djdoeslinux/choreobot/bot"
	"github.com/djdoeslinux/choreobot/command/counter"
	"github.com/djdoeslinux/choreobot/command/loyalty_points"
	"github.com/djdoeslinux/choreobot/command/turing_test"
	"github.com/djdoeslinux/choreobot/meter"
	"github.com/djdoeslinux/choreobot/moderator"
	"github.com/djdoeslinux/choreobot/straw_poll"
	"github.com/djdoeslinux/choreobot/user"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	var models []interface{}
	models = append(models, bot.Models...)
	models = append(models, counter.Models...)
	models = append(models, loyalty_points.Models...)
	models = append(models, meter.Models...)
	models = append(models, moderator.Models...)
	models = append(models, straw_poll.Models...)
	models = append(models, turing_test.Models...)
	models = append(models, user.Models...)

	db.AutoMigrate(models...)
}

type Wrapper interface {
}

type Implementor interface {
}

type Event interface {
}

type Response interface {
}
