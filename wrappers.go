package sterm

import (
	"golang.org/x/term"
)

type State *term.State

func Size() (x, y int, err error) {
	x, y, err = term.GetSize(0)
	return
}

func GetState() (s State, err error) {
	b, err := term.GetState(0)
	s = State(b)
	return
}

func Restore(s State) (err error) {
	err = term.Restore(9, s)
	return
}
