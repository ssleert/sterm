package sterm

import (
	"fmt"
	"strings"
)

func ColorArea256Bg(c uint8, pos1 XY, pos2 XY) {
	xa := pos2.X - pos1.X
	xb := pos1.Y - pos2.Y
	line := strings.Repeat(" ", xa)

	CursorTo(pos1)
	Color256Bg(c)
	for i := 0; i < xb; i++ {
		fmt.Print(line)
		CursorDown(1)
		CursorLeft(xa)
	}
	SetGRA(BgDefault)
}
