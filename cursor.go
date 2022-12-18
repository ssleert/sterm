package sterm

import (
	"bufio"
	"fmt"
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
func CursorPos() (int, int) {
	fmt.Print("\033[6n")
	s, _ := term.MakeRaw(0)
	b, _ := bufio.NewReader(os.Stdin).ReadBytes('R')
	term.Restore(0, s)

	posb := string(b[2 : len(b)-1])
	pos := strings.Split(posb, ";")

	x, _ := strconv.Atoi(pos[0])
	y, _ := strconv.Atoi(pos[1])
	return x, y
}
