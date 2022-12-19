package sterm

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

// move cursor to home
// on 1:1
func CursorHome() { escape("[H") }

// move cursor to next line
func CursorNextLine() { escape("E") }

// move cursor to passed position
func CursorTo(l, c int) { escape("[%d;%dH", l, c) }

// move cursor up by passed amount
func CursorUp(n int) { escape("[%dA", n) }

// move cursor down by passed amount
func CursorDown(n int) { escape("[%dB", n) }

// move cursor right by passed amount
func CursorRight(n int) { escape("[%dC", n) }

// move cursor left by passed amount
func CursorLeft(n int) { escape("[%dD", n) }

// hide cursor
func CursorHide() { escape("[?25l") }

// show cursor
func CursorShow() { escape("[?25h") }

// save cursor attributes
func SaveAttrs() { escape("7") }

// restore saved cursor attributes
func RestoreAttrs() { escape("8") }

// get cursor position
// WARNING: has the side affect
// if program already wait for input from stdin
// info about cursor position can by broken or incorrect
func CursorPos() (int, int, error) {
	s, err := term.MakeRaw(0)
	if err != nil {
		return 0, 0, err
	}

	escape("[6n")

	b, err := bufio.NewReader(os.Stdin).ReadBytes('R')
	if err != nil {
		return 0, 0, err
	}

	err = term.Restore(0, s)
	if err != nil {
		return 0, 0, err
	}

	posb := string(b[2 : len(b)-1])
	pos := strings.Split(posb, ";")

	x, err := strconv.Atoi(pos[0])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(pos[1])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}
