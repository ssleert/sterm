package sterm

import (
	"golang.org/x/term"
)

// simple wrapper for term.GetSize()
// to remove unused imports)
// all docs https://pkg.go.dev/golang.org/x/crypto/ssh/terminal#GetSize
func Size() (int, int, error) {
	x, y, err := term.GetSize(0)
	return x, y, err
}

