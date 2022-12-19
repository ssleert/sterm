package sterm

import "fmt"

// set graphics rendition attributes for terminal
func SetGRA(gras ...string) {
	if len(gras) == 0 {
		return
	}

	for _, s := range gras {
		fmt.Print(s)
	}
}

// set 256 foreground color
func Color256Fg(c int) { escape("[38;5;%dm", c) }

// set 256 background color
func Color256Bg(c int) { escape("[48;5;%dm", c) }

// set RGB foreground color
func ColorRGBFg(r, g, b int) { escape("[38;2;%d;%d;%dm", r, g, b) }

// set RGB background color
func ColorRGBBg(r, g, b int) { escape("[48;2;%d;%d;%dm", r, g, b) }
