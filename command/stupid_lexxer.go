package command

import (
	"github.com/gempir/go-twitch-irc"
	"strings"
)

type stupidLexer struct {
	tokens     []string
	curToken   int
	tokenCount int
	raw        twitch.Message
}

func newStupidLexxer(m twitch.Message) TokenStream {
	trimmed := strings.TrimSpace(m.Text)
	tokens := strings.Split(trimmed, " ")
	return &stupidLexer{tokens: tokens, curToken: 1, tokenCount: len(tokens), raw: m}
}

func (s *stupidLexer) GetCommand() string {
	return s.tokens[0]
}

func (s *stupidLexer) NumArgs() int {
	return len(s.tokens) - 2
}

func (s *stupidLexer) GetToken(index int) Token {
	i := index + 1
	return Token(s.tokens[i])
}

func (s *stupidLexer) NotDone() bool {
	return s.curToken < s.tokenCount
}

func (s *stupidLexer) Next() {
	s.curToken++
}

func (s *stupidLexer) CurrentToken() Token {
	return Token(s.tokens[s.curToken])
}

func (s *stupidLexer) GetRaw() twitch.Message {
	return s.raw
}
