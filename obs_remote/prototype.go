/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */
package obs_remote

import (
	"encoding/json"
	"fmt"
	"github.com/djdoeslinux/choreobot/client"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/gempir/go-twitch-irc"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/url"
	"time"
)

var c *websocket.Conn

var LittlePrimary SceneItem
var BigPartner SceneItem
var MavFeedText SceneItem
var KinFeedText SceneItem
var PreviewScene SceneItem
var LittleTVMute SetMute
var BigTVMute SetMute

var Mav command.Command
var Kin command.Command
var TogglePreview command.Command

func init() {
	u := url.URL{Scheme: "ws", Host: "10.0.0.17:4444"}
	c, _, _ = websocket.DefaultDialer.Dial(u.String(), nil)
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer c.Close()
		for {
			_, m, e := c.ReadMessage()
			if e != nil {
				fmt.Println("read: ", e)
			}
			fmt.Println(string(m))
		}
	}()
	LittlePrimary = SceneItem{Name: "SetSceneItemProperties", TargetScene: "GameOnlyCutscene", ItemName: "LittleTV", Visible: false}
	BigPartner = SceneItem{Name: "SetSceneItemProperties", TargetScene: "PartnerGame", ItemName: "BigTV", Visible: false}
	MavFeedText = SceneItem{Name: "SetSceneItemProperties", TargetScene: "PartnerGame", ItemName: "MavIsPartnerText", Visible: false}
	KinFeedText = SceneItem{Name: "SetSceneItemProperties", TargetScene: "PartnerGame", ItemName: "KinIsPartnerText", Visible: false}
	PreviewScene = SceneItem{Name: "SetSceneItemProperties", TargetScene: "Duos Cam Right", ItemName: "PartnerGame", Visible: true}
	LittleTVMute = SetMute{Method: "SetMute", Name: "LittleTV", Muted: false}
	BigTVMute = SetMute{Method: "SetMute", Name: "BigTV", Muted: false}
	Mav = &com{"mav", WatchMav}
	Kin = &com{"kin", WatchKin}
	TogglePreview = &com{"previewOn", Preview}
}

type com struct {
	focus string
	exe   func(event *client.TwitchEvent, stream command.TokenStream) command.Result
}

func (*com) IsAllowed(u twitch.User) bool {
	return true
}

func (c *com) Evaluate(e *client.TwitchEvent, t command.TokenStream) command.Result {
	return c.exe(e, t)
}

type Req struct {
	Name string `json:"request-type"`
	ID   string `json:"message-id"`
}

type SetScene struct {
	Name        string `json:"request-type"`
	ID          string `json:"message-id"`
	TargetScene string `json:"scene-name"`
}

type SceneItem struct {
	Name        string `json:"request-type"`
	ID          string `json:"message-id"`
	TargetScene string `json:"scene-name"`
	ItemName    string `json:"item"`
	//Item string `json:"name"`
	Visible bool `json:"visible"`
	//Rotation float32 `json:"rotation"`
	//Pos `json:"position"`
	//Bounds `json:"bounds"`
}

type SetMute struct {
	Method string `json:"request-type"`
	ID     string `json:"message-id"`
	Name   string `json:"source"`
	Muted  bool   `json:"mute"`
}

type Pos struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Alignment string `json:"alignment"`
}

type Bounds struct {
	BoundType string `json:"type"`
}

type Crop struct {
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
	Right  int `json:"Right"`
	Top    int `json:"Top"`
}

type Scale struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func WatchMav(event *client.TwitchEvent, stream command.TokenStream) command.Result {
	LittlePrimary.Visible = false
	BigPartner.Visible = false
	BigTVMute.Muted = false
	LittleTVMute.Muted = true
	//KinFeedText.Visible = false
	MavFeedText.Visible = false
	setVis()
	return command.TODO
}

func Preview(event *client.TwitchEvent, stream command.TokenStream) command.Result {
	PreviewScene.Visible = !PreviewScene.Visible
	c.WriteJSON(PreviewScene)
	return command.TODO
}

func WatchKin(event *client.TwitchEvent, stream command.TokenStream) command.Result {
	LittlePrimary.Visible = true
	LittleTVMute.Muted = false
	BigTVMute.Muted = true
	BigPartner.Visible = true
	//KinFeedText.Visible = true
	MavFeedText.Visible = true
	setVis()
	return command.TODO
}

func setVis() {
	c.WriteJSON(LittlePrimary)
	c.WriteJSON(BigPartner)
	c.WriteJSON(MavFeedText)
	c.WriteJSON(LittleTVMute)
	c.WriteJSON(BigTVMute)
	//c.WriteJSON(KinFeedText)
}

func doit() {
	//gv := Req{Name: "GetVersion", ID: uuid.NewV4().String()}
	ls := Req{Name: "GetSceneList", ID: uuid.NewV4().String()}
	//switchToCutscene := SetScene{Name: "SetCurrentScene", ID: "switch to cutscene", TargetScene: "GameOnlyCutscene"}
	//switchToDuos := SetScene{Name: "SetCurrentScene", ID: "switch to cutscene", TargetScene: "Duos Cam Right"}
	getItem := SceneItem{Name: "GetSceneItemProperties", ID: "Get items", TargetScene: "PartnerGame", ItemName: "PartnetGameText"}
	setLittle := SceneItem{Name: "SetSceneItemProperties", ID: "Set little on", TargetScene: "GameOnlyCutscene", ItemName: "LittleTV", Visible: true}

	//dumpjs(setLittle)
	//c.WriteJSON(gv)
	//time.Sleep(500 * time.Millisecond)
	c.WriteJSON(ls)
	//time.Sleep(500 * time.Millisecond)
	c.WriteJSON(getItem)
	time.Sleep(5000 * time.Millisecond)
	//c.WriteJSON(switchToCutscene)
	//time.Sleep(2000 * time.Millisecond)

	setLittle.Visible = true
	c.WriteJSON(setLittle)
	time.Sleep(3000 * time.Millisecond)

	setLittle.Visible = false
	setLittle.ID = "Set little off"
	time.Sleep(5 * time.Second)
	c.WriteJSON(setLittle)
	//c.WriteJSON(switchToDuos)
	time.Sleep(6000 * time.Second)
	return
	//json := `{"request-type": ""`
	//c.WriteMessage(websocket.TextMessage, )
}

func dumpjs(x interface{}) {
	s, _ := json.Marshal(x)
	fmt.Println(string(s))
}
