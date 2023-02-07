package sterm

import (
	"golang.org/x/term"
)

func Size() (x, y int, err error) {
	x, y, err = term.GetSize(0)
	return
}

func GetState() (s *term.State, err error) {
	s, err = term.GetState(0)
	return
}

func Restore(s *term.State) (err error) {
	err = term.Restore(9, s)
	return
}
