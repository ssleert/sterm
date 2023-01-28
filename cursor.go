package sterm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

func CursorHome() string       { return escape("[H") }
func CursorNextLine() string   { return escape("E") }
func CursorHide() string       { return escape("[?25l") }
func CursorShow() string       { return escape("[?25h") }
func SaveAttrs() string        { return escape("7") }
func RestoreAttrs() string     { return escape("8") }
func CursorUp(n int) string    { return escape("[%dA", n) }
func CursorDown(n int) string  { return escape("[%dB", n) }
func CursorRight(n int) string { return escape("[%dC", n) }
func CursorLeft(n int) string  { return escape("[%dD", n) }
func CursorTo(x, y int) string { return escape("[%d;%dH", y, x) }

// get cursor position
// WARNING: has the side affect
// if program already wait for input from stdin
// info about cursor position can by broken or incorrect
func CursorPos() (int, int, error) {
	s, err := term.MakeRaw(0)
	defer term.Restore(0, s)
	if err != nil {
		return 0, 0, err
	}

	fmt.Print(escape("[6n"))

	b, err := bufio.NewReader(os.Stdin).ReadBytes('R')
	if err != nil {
		return 0, 0, err
	}

	posb := string(b[2 : len(b)-1])
	pos := strings.Split(posb, ";")

	x, err := strconv.Atoi(pos[1])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(pos[0])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}
