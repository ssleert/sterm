package sterm

import "strings"

func ClearLineRight() string    { return escape("[0K") }
func ClearLineLeft() string     { return escape("[1K") }
func ClearEntireLine() string   { return escape("[2K") }
func ClearScreenDown() string   { return escape("[0J") }
func ClearScreenUp() string     { return escape("[1J") }
func ClearEntireScreen() string { return escape("[2J") }

// clear screen and move cursor to 1:1
func ClearHome() string {
	var s strings.Builder
	s.Grow(6)
	s.WriteString(ClearEntireScreen())
	s.WriteString(CursorHome())
	return s.String()
}
