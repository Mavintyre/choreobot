package command

import (
	"github.com/djdoeslinux/choreobot/x"
	"github.com/gempir/go-twitch-irc"
	"strings"
)

type stupidLexer struct {
	tokens     []Token
	curToken   int
	tokenCount int
	raw        twitch.Message
}

type stupidToken struct {
	v string
}

func (s *stupidToken) String() string {
	return s.v
}

func (s *stupidLexer) NumTokens() int {
	return s.tokenCount
}

func (s *stupidLexer) GetTokenByIndex(index int) Token {
	if index > s.tokenCount || index < 0 {
		return nil
	}
	return s.tokens[index]
}

func (s *stupidLexer) Prev() error {
	if s.curToken > 0 {
		s.curToken--
		return nil
	}
	return x.OutOfBounds
}

func (s *stupidLexer) Token() Token {
	return s.tokens[s.curToken]
}

func (s *stupidLexer) PopToken() (Token, error) {
	panic("implement me")
}

func newStupidLexxer(m twitch.Message) TokenStream {
	trimmed := strings.TrimSpace(m.Text)
	var tok []Token
	for _, t := range strings.Split(trimmed, " ") {
		tok = append(tok, Token(&stupidToken{t}))
	}

	return &stupidLexer{tokens: tok, curToken: 1, tokenCount: len(tok), raw: m}
}

func (s *stupidLexer) GetCommand() Token {
	return s.tokens[0]
}

func (s *stupidLexer) NumArgs() int {
	return len(s.tokens) - 2
}

func (s *stupidLexer) GetToken(index int) Token {
	i := index + 1
	return s.tokens[i]
}

func (s *stupidLexer) NotDone() bool {
	return s.curToken < s.tokenCount
}

func (s *stupidLexer) Next() error {
	s.curToken++
	return nil
}

func (s *stupidLexer) CurrentToken() Token {
	return Token(s.tokens[s.curToken])
}

func (s *stupidLexer) GetRaw() twitch.Message {
	return s.raw
}

func (s *stupidLexer) Seek(i int) error {
	if s.tokenCount < i {
		return x.OutOfBounds
	}
	s.curToken = i
}
