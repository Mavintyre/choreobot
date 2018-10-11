/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package command

import (
	"fmt"
	"testing"
)

type test struct {
	in, out, title string
	index int
}

var tests []test = []test{
test{`!addQuote This is a simple text string broken on whitespace`, "string", "Simple string correct", 6},
test{`!addTuring "This is a " quoted string`, "This is a ", "Quoted string correct", 1},
test{`!addTuring This is a { @templated interpolation }`, " @templated interpolation ", "Simple paired correct", 4},
test{`!addTuring This is a ( {{template}} that is embedded ) "With some " quotes`, " {{template}} that is embedded ", "Multiple pairing characters", 4},
test{`!addTuring {a{{bc}} de {f} g}`, "a{{bc}} de {f} g", "Nested pairings", 1},
}

func Test_Cases(t *testing.T) {
	for i, x := range tests {
		fmt.Println("Running test: ", i)
		_, c := lex(x.in)
		s := sliceify(c)
		assertEqual(t, s[x.index].String(), x.out, "")
	}
}

func sliceify(c chan Token) (s []Token){
	for t := range c {
		s = append(s, t)
	}
	return
}

func assertEqual(t *testing.T, a, b, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}