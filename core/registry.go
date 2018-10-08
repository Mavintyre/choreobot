/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package core

import (
	"github.com/djdoeslinux/choreobot/bot"
	"github.com/djdoeslinux/choreobot/user"
	"github.com/jinzhu/gorm"
)

var models map[Namespace][]interface{}

type WrapperFactory func(user.UserI) Wrapper
type IdentityFactory func(user.User) user.UserI
type ImplementorFactory func(bot.Bot) Implementor

var identityProviders map[Namespace]IdentityFactory
var wrappers map[Namespace]WrapperFactory
var implementors map[Namespace]ImplementorFactory

func init() {
	models = make(map[Namespace][]interface{})
	identityProviders = make(map[Namespace]IdentityFactory)
	wrappers = make(map[Namespace]WrapperFactory)
	implementors = make(map[Namespace]ImplementorFactory)
}

func RegisterWrapper(name Namespace, factory func(user.UserI) Wrapper){

}

func RegisterIdentifier(name Namespace, factory func(user.User) user.UserI){

}

func RegisterImplementor(name Namespace, factory func(bot bot.Bot) Implementor){

}



func RegisterModel(name Namespace, modelsToRegister ...interface{}) {
	models[name] = append(models[name], modelsToRegister...)
}

func GetUserForNamespace(name Namespace, u user.User) user.UserI{
	if f, exists := identityProviders[name]; exists {
		return f(u)
	}
	return nil
}

func Get


func AutoMigrate(db *gorm.DB) {

	for _, list := range models {
		db.AutoMigrate(list...)
	}
}

