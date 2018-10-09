/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package command

import (
	"fmt"
	"github.com/djdoeslinux/choreobot/client"
	"github.com/gempir/go-twitch-irc"
	"unicode/utf8"
)

type lexFn func(*lexer) lexFn

type tokenType int

const (
	_ tokenType = iota
	simpleText
	quotedText

)

const eof rune = 0

var closerOf map[rune]rune

const escape rune = '\\'

func init(){
	closerOf = make(map[rune]rune)
	closerOf['\''] = '\''
	closerOf['{'] = '}'
	closerOf['['] = ']'
	closerOf['"'] = '"'
	closerOf['<'] = '>'
}



type properToken struct {
	tokenType
	value string
}

func (*properToken) String() string {
	panic("implement me")
}


type lexedStream struct {
	tokens []Token
	event *client.TwitchEvent
	numTokens int
	pos int
}

func LexTwitchEvent(t *client.TwitchEvent) TokenStream{
	_, c := lex(t)
	s := &lexedStream{}
	for tok := range c{
		s.tokens = append(s.tokens, tok)
		s.numTokens++
	}
	return s
}

func (l *lexedStream) GetRaw() twitch.Message {
	return *l.event.Message
}

func (l *lexedStream) NumTokens() int {
	return l.numTokens
}

func (l *lexedStream) GetTokenByIndex(index int) Token {
	if index > 0 && index < l.numTokens{
		return l.tokens[index]
	}
	return nil
}

func (l *lexedStream) NotDone() bool {
	panic("implement me")
}

func (l *lexedStream) Next() error {
	panic("implement me")
}

func (l *lexedStream) Prev() error {
	panic("implement me")
}

func (l *lexedStream) Token() Token {
	panic("implement me")
}

func (l *lexedStream) PopToken() (Token, error) {
	panic("implement me")
}

func (l *lexedStream) Seek(i int) error {
	l.pos = i
}

type lexer struct {
	emitter chan Token
	input string
	start int
	pos int
	width int
	nextCloser rune
	event *client.TwitchEvent
}

func lex(t *client.TwitchEvent) (*lexer, chan Token){
	l := &lexer{event: t, input: t.Message.Text}
	l.emitter = make(chan Token)
	go l.run()
	return l, l.emitter
}

func (l *lexer) run() {
	for state := rootState; state != nil; {
		state = state(l)
	}
	close(l.emitter)
}

func (l *lexer) emit(t tokenType){
	l.emitter <- &properToken{
		tokenType: t,
		value: l.input[l.start:l.pos],
	}
	l.start = l.pos
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width =
		utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return
}

func rootState(l *lexer) lexFn{
	for {
		r := l.next()
		switch r {
		case escape:
			l.next()
		case eof:
			break
		case ' ', '\t':
			l.start = l.pos
		default:
			if o, exists := closerOf[r]; exists{
				l.nextCloser = o
				return findCloser
			}
		}
	}
	if l.pos > l.start {
		l.emit(simpleText)
	}
	return nil
}

func findCloser(l *lexer) lexFn{
	for {
		r := l.next()
		switch r{
		case eof:
			fmt.Println("Error parsing enclosed string. Did not find ", l.nextCloser , " before end of string.")
			return nil
		case escape:
			l.next()
		case l.nextCloser:
			l.emit(quotedText)
			return rootState
		}
	}
}