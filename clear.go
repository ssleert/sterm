package sterm

import "fmt"

func ClearLineRight() string    { return escape("[0K") }
func ClearLineLeft() string     { return escape("[1K") }
func ClearEntireLine() string   { return escape("[2K") }
func ClearScreenDown() string   { return escape("[0J") }
func ClearScreenUp() string     { return escape("[1J") }
func ClearEntireScreen() string { return escape("[2J") }

// clear screen and move cursor to 1:1
func ClearHome() {
	fmt.Print(ClearEntireScreen())
	fmt.Print(CursorHome())
}
