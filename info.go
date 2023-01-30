package sterm

import (
	"golang.org/x/term"
)

// simple wrapper for term.GetSize()
// for code minimization
// all docs https://pkg.go.dev/golang.org/x/crypto/ssh/terminal#GetSize
func Size() (int, int, error) {
	x, y, err := term.GetSize(0)
	return x, y, err
}
