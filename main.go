/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package main

import (
	"github.com/djdoeslinux/choreobot/bot"
	"github.com/djdoeslinux/choreobot/core"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sanity-io/litter"
)

//Broad TODOs
// Choose a persistence approach
func main() {
	//Just hardcode a sqlite path for now so we can
	db, _ := gorm.Open("sqlite3", "/tmp/choreobot.sqlite")
	defer db.Close()
	var b []*bot.Bot
	core.AutoMigrate(db)

	//db.Create(&core.Bot{ UserName:   "stupidbot", OAuthToken: "", })
	db.Set("gorm:auto_preload", true).Find(&b)
	for _, bot := range b {
		bot.Start(db)
		litter.Dump(bot)
	}
}
