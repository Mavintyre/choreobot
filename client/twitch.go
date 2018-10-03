package client

import (
	"fmt"
	"github.com/gempir/go-twitch-irc"
	"os"
	"time"
)

var client *twitch.Client

func connect() {
	oauthToken := os.Getenv("CHOREOBOT_TOKEN")
	client = twitch.NewClient("choreobot", fmt.Sprintf("oauth:%s", oauthToken))
	client.OnConnect(handleOnConnect)
	client.OnNewWhisper(handleOnNewWhisper)
	client.OnNewMessage(handleOnNewMessage)
	client.OnNewRoomstateMessage(handleOnNewRoomstateMessage)
	client.OnNewClearchatMessage(handleOnNewClearchatMessage)
	client.OnNewUsernoticeMessage(handleOnNewUsernoticeMessage)
	client.OnNewUserstateMessage(handleOnNewUserstateMessage)
	client.Join("djdoeslinux")
	err := client.Connect()
	if err != nil {
		fmt.Println(err)
	}
}

func handleOnConnect() {
	msg := fmt.Sprintf("I joined today at %s", time.Now())
	client.Say("djdoeslinux", msg)

}

func handleOnNewWhisper(user twitch.User, message twitch.Message) {

}
func handleOnNewMessage(channel string, user twitch.User, message twitch.Message) {
	fmt.Printf("At %s %s said: %s\n", message.Time, user.Username, message.Text)
}
func handleOnNewRoomstateMessage(channel string, user twitch.User, message twitch.Message) {

}
func handleOnNewClearchatMessage(channel string, user twitch.User, message twitch.Message) {

}
func handleOnNewUsernoticeMessage(channel string, user twitch.User, message twitch.Message) {

}
func handleOnNewUserstateMessage(channel string, user twitch.User, message twitch.Message) {

}
