/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package turing_test

import (
	"fmt"
	"github.com/djdoeslinux/choreobot/client"
	"github.com/djdoeslinux/choreobot/command"
	"github.com/djdoeslinux/choreobot/user"
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

const maxRandomTries int = 5

var Models []interface{}

func init() {
	Models = append(Models, &Turing{}, &Response{})
}

type Turing struct {
	gorm.Model
	Name           string
	EventHandlerID uint
	Responses      []Response

	//Private Members
	numResponses     int
	responsesByIndex []*Response
	isInitialized    bool
}

type Response struct {
	gorm.Model
	TuringID uint
	Index    int
	Template string
	// Private Members
	compiled         *template.Template
	hasTemplateError bool
}

type Responder interface {
	command.Command
	GetResponseByIndex(i int) string
	GetRandomResponse() string
	AddResponse(db *gorm.DB, template string) int
}

func GetTuringByEventHandlerID(db *gorm.DB, eventHandlerID uint) (t *Turing) {
	t = &Turing{EventHandlerID: eventHandlerID}
	db.Find(t)
	t.initialize()
	return
}

func NewBlankTuring() (t *Turing) {
	t = &Turing{}
	return t
}

func NewTuring(db *gorm.DB, name string, eventHandlerID uint, responses ...string) (*Turing, error) {
	t := &Turing{Name: name, EventHandlerID: eventHandlerID}
	t.numResponses = len(responses)
	db.Create(t).Find(t)
	for _, templ := range responses {
		t.AddResponse(db, templ)
	}
	db.Save(t)
	return t, nil
}

func (t *Turing) AddResponse(db *gorm.DB, templateString string) error {
	t.numResponses++
	index := t.numResponses
	r := Response{TuringID: t.ID, Index: index, Template: templateString}
	db.Save(&r)
	return t.initResponse(r)
}

func (t *Turing) ModifyResponse(db *gorm.DB, responseIndex int, newTemplate string) error {
	r := t.responsesByIndex[responseIndex]
	err := r.modify(db, newTemplate)
	if err != nil {
		return err
	}
	// Need to refresh self from the database to pickup changes in t.Responses
	db.Find(t)
	t.initialize()
	return nil
}

func (t *Turing) Save(db *gorm.DB) {
	db.Save(t)
}

func (r *Response) modify(db *gorm.DB, newTemplate string) error {
	old := r.Template
	oldCompile := r.compiled
	r.Template = newTemplate
	err := r.compile()
	if err != nil {
		r.Template = old
		r.compiled = oldCompile
		return err
	}
	db.Save(r)
	return nil
}

func (t *Turing) initialize() {

	for _, r := range t.Responses {
		t.initResponse(r)
	}
	t.isInitialized = true
}

func (t *Turing) initResponse(r Response) error {
	// Verify our slice is big enough for the configured index.
	if len(t.responsesByIndex) < r.Index {
		newSlice := make([]*Response, r.Index, r.Index)
		copy(newSlice, t.responsesByIndex)
		t.responsesByIndex = newSlice
	}
	t.responsesByIndex[r.Index] = &r
	return r.compile()
}

func (r Response) compile() (err error) {
	templateName := strconv.Itoa(int(r.ID))
	templ := template.New(templateName)
	r.compiled, err = templ.Parse(r.Template)
	//Error passed through silently if existing
	//TODO: Log the error
	return
}

func (t *Turing) Evaluate(e *client.TwitchEvent, args command.TokenStream) command.Result {
	if t.numResponses == 1 {
		return t.doTemplate(0, e, args)
	}
	if args.NumTokens() == 1 {
		return t.doRandomTemplate(e, args)
	}

	index, err := strconv.Atoi(args.GetTokenByIndex(1).String())
	if err == nil {
		return t.doTemplate(index, e, args)
	}
	return command.Error(err)
}

func (t *Turing) doRandomTemplate(e *client.TwitchEvent, s command.TokenStream) command.Result {
	startIndex := rand.Intn(t.numResponses)
	for i := 0; i < maxRandomTries; i++ {
		index := (startIndex + i) % t.numResponses
		r, err := t.responsesByIndex[index].doTemplate(e, s)
		if err != nil {
			return r
		}
	}
	return command.Error(fmt.Errorf("No good templates found after %s tries. Giving up", maxRandomTries))
}

func (t *Turing) doTemplate(i int, e *client.TwitchEvent, stream command.TokenStream) command.Result {
	r, _ := t.responsesByIndex[i].doTemplate(e, stream)
	return r
}

func (r *Response) doTemplate(e *client.TwitchEvent, s command.TokenStream) (command.Result, error) {
	if r.hasTemplateError {
		return nil, fmt.Errorf("Malformed Response Template.")
	}
	data := struct {
		me   *user.User
		args []string
	}{}

	s.Seek(1)
	for s.NotDone() {
		token, _ := s.PopToken()
		data.args = append(data.args, token.String())
	}
	data.me = user.GetUserByEvent(e)
	result := strings.Builder{}
	err := r.compiled.Execute(&result, data)
	if err != nil {
		return command.Error(err), err
	}
	return command.TODO, nil
}
