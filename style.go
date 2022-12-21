package sterm

import (
	"fmt"
)

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
func Color256Fg(c uint8) { escape("[38;5;%dm", c) }

// set 256 background color
func Color256Bg(c uint8) { escape("[48;5;%dm", c) }

// set RGB foreground color
func ColorRGBFg(c RGB) { escape("[38;2;%d;%d;%dm", c.R, c.G, c.B) }

// set RGB background color
func ColorRGBBg(c RGB) { escape("[48;2;%d;%d;%dm", c.R, c.G, c.B) }
