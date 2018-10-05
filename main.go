/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package main

import (
	"github.com/djdoeslinux/choreobot/core"
	"github.com/djdoeslinux/choreobot/registry"
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
	var b []*core.Bot
	registry.AutoMigrate(db)

	//db.Create(&core.Bot{ UserName:   "stupidbot", OAuthToken: "", })
	db.Find(&b)
	litter.Dump(b)
}
