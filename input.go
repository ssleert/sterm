package sterm

import (
	"bufio"
	"os"

	"golang.org/x/term"
)

func GetChar() (rune, error) {
	s, err := term.MakeRaw(0)
	defer term.Restore(0, s)
	if err != nil {
		return 0, err
	}

	in := bufio.NewReader(os.Stdin)

	rn, _, err := in.ReadRune()
	if err != nil {
		return 0, err
	}

	return rn, nil
}
