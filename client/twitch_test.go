package client

import (
	"fmt"
	"github.com/gempir/go-twitch-irc"
	"github.com/sanity-io/litter"
	"os"
	"testing"
	"time"
)

func Test_Connect(t *testing.T) {
	d := &debugUser{}
	d.client = connect(d)
}

type debugUser struct {
	client *twitch.Client
}

func (d *debugUser) GetUserName() string {
	return "djdoeslinux"
}

func (d *debugUser) GetToken() string {
	return os.Getenv("CHOREOBOT_TOKEN")
}

func (d *debugUser) GetChannelList() []string {
	return []string{"djdoeslinux"}
}

func (d *debugUser) OnConnect() {
	msg := fmt.Sprintf("I joined today at %s", time.Now())
	fmt.Println(msg)
	//client.Say("djdoeslinux", msg)
}

func (d *debugUser) OnWhisper(user twitch.User, message twitch.Message) {
	fmt.Println("Whisper")
	litter.Dump(user)
	litter.Dump(message)
}

func (d *debugUser) OnMessage(channel string, user twitch.User, message twitch.Message) {
	fmt.Printf("In channel %s At %s %s said: %s\n", channel, format(message.Time), user.Username, message.Text)
	litter.Dump(channel)
	litter.Dump(user)
	litter.Dump(message)
}

func (d *debugUser) OnRoomState(channel string, user twitch.User, message twitch.Message) {
	fmt.Printf("Channel %s just changed roomstate at %s because user %s said %s\n", channel, format(message.Time), user.Username, message.Raw)
	litter.Dump(channel)
	litter.Dump(user)
	litter.Dump(message)
}

func (d *debugUser) OnClearChat(channel string, user twitch.User, message twitch.Message) {
	fmt.Println("Clear chat")
	litter.Dump(channel)
	litter.Dump(user)
	litter.Dump(message)
}

func (d *debugUser) OnUserNotice(channel string, user twitch.User, message twitch.Message) {
	litter.Dump(channel)
	litter.Dump(user)
	litter.Dump(message)
}

func (d *debugUser) OnUserState(channel string, user twitch.User, message twitch.Message) {
	litter.Dump(channel)
	litter.Dump(user)
	litter.Dump(message)
}

func format(t time.Time) string {
	return t.Format(time.Kitchen)
}
