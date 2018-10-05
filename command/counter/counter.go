/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package counter

import "github.com/jinzhu/gorm"

// This will implement a counter, with goals and potentially automatic triggers based on those goals.
// Manual Counter: !addswear -- increments the counter, reports the current swear count, and gives everyone 50 musk.
// Dynamic Counter: !uptime -- how long since the stream last started
var Models []interface{}

func init() {
	Models = append(Models, &Counter{})
}

type Counter struct {
	gorm.Model
	ChannelID uint
	Name      string
	Count     int
}
