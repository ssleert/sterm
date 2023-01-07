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
func Color256Fg(c int) error {
	if c < 0 || c > 255 {
		return ErrColorOutOfRange
	}
	escape("[38;5;%dm", c)
	return nil
}

// set 256 background color
func Color256Bg(c int) error {
	if c < 0 || c > 255 {
		return ErrColorOutOfRange
	}
	escape("[48;5;%dm", c)
	return nil
}

// set RGB foreground color
func ColorRGBFg(c RGB) { escape("[38;2;%d;%d;%dm", c.R, c.G, c.B) }
func ColorRGBFgStr(c RGB) string {
	return fmt.Sprintf("[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

// set RGB background color
func ColorRGB(c RGB) { escape("[48;2;%d;%d;%dm", c.R, c.G, c.B) }
func ColorRGBBgStr(c RGB) string {
	return fmt.Sprintf("[38;2;%d;%d;%dm", c.R, c.G, c.B)
}
