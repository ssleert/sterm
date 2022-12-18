package sterm

import (
	"bufio"
	"os"

	"golang.org/x/term"
)

// simple func to get char from terminal
// similar to getchar() in C
func GetChar() rune {
	s, _ := term.MakeRaw(0)
	in := bufio.NewReader(os.Stdin)
	rn, _, _ := in.ReadRune()
	term.Restore(0, s)
	return rn
}
