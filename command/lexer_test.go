/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package command

import (
	"fmt"
	"testing"
)

var normal string = `!addQuote This is a simple text string broken on whitespace`
var quoted string = `!addTuring "This is a " quoted string`
var paired string = `!addTuring This is a { @templated interpolation }`
var embeddedPaired string = `!addTuring This is a ( {{template}} that is embedded ) "With some " quotes`

func Test_Normal(t *testing.T) {
	_, c := lex(normal)
	s := sliceify(c)
	assertEqual(t, s[6].String(), "string", "string is string")
}

func Test_Quoted(t *testing.T) {
	_, c := lex(quoted)
	s := sliceify(c)
	assertEqual(t, s[1].String(), "This is a ", "quoted string is string")
}

func Test_Paired(t *testing.T) {
	_, c := lex(paired)
	s := sliceify(c)
	assertEqual(t, s[4].String(), " @templated interpolation ", "template clean")
}

func sliceify(c chan Token) (s []Token){
	for t := range c {
		s = append(s, t)
	}
	return
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}