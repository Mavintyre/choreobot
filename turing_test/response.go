package turing_test

import (
	"fmt"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/gempir/go-twitch-irc"
	"github.com/jinzhu/gorm"
	"math/rand"
	"strconv"
	"strings"
	"text/template"
)

//This is a generic responder command implementation.
// It can handle any rote stateless responses by template.
// Use cases:
// Simple Response: !rules -- Just says something verbatim in the chat (simple is actually just a 1 template multiresponse)
// MultiResponse: !quote [int|tag] -- Keeps track of multiple possible responses. Picks randomly unless a specific index is requested
///// MultiResponse
// Templated: !law {U} {1} {2} {...} -- Interpolates a response based on a template. Arguments are space separated. An ellipsis indicates "all remaining arguments"

type Turing struct {
	gorm.Model
	Name      string
	ChannelID uint
	Responses []Response

	//Private Members
	numResponses     int
	templatesByIndex []*template.Template
	isInitialized    bool
}

type Response struct {
	gorm.Model
	TuringID uint
	Index    int
	Template string
}

type Responder interface {
	command.Command
	GetResponseByIndex(i int) string
	GetRandomResponse() string
	AddResponse(string) int
}

func GetTuringByChannelAndName(channelID uint, name string) *Turing {
	panic("x")
}

func (t *Turing) initialize() {

	for _, r := range t.Responses {
		if len(t.templatesByIndex) < r.Index {
			newSlice := make([]*template.Template, r.Index, r.Index)
			copy(newSlice, t.templatesByIndex)
			t.templatesByIndex = newSlice
		}
		templ := template.New(fmt.Sprintf("%s:%s", t.Name, strconv.Itoa(r.Index)))
		parsedTemplate, err := templ.Parse(r.Template)
		if err != nil {
			panic("broken template?")
		}
		t.templatesByIndex[r.Index] = parsedTemplate
	}
}

func (t *Turing) Evaluate(u twitch.User, args command.TokenStream) command.Result {
	if t.numResponses == 1 {
		return t.doTemplate(0, u, args)
	}
	if args.NumTokens() == 1 {
		return t.doTemplate(rand.Intn(t.numResponses), u, args)
	}

	index, err := strconv.Atoi(args.GetTokenByIndex(1).String())
	if err == nil {
		return t.doTemplate(index, u, args)
	}
	return command.Error(err)
}

func (t *Turing) AddResponse(string) int {
	panic("implement me")
}

func (t *Turing) doTemplate(i int, user twitch.User, stream command.TokenStream) command.Result {
	tt := t.templatesByIndex[i]
	data := struct {
		me   string
		args []string
	}{}
	stream.Seek(i)
	for stream.NotDone() {
		token, _ := stream.PopToken()
		data.args = append(data.args, token.String())
	}
	data.me = user.Username
	result := strings.Builder{}
	err := tt.Execute(&result, data)
	if err != nil {
		return command.Error(err)
	}
	return command.TODO
}
