package sterm

import (
	"errors"
	"fmt"
	"strings"
)

// print some characters in defined area
func CharArea(ch rune, pos1 XY, pos2 XY) error {
	if pos1 == pos2 {
		return errors.New("the positions of the points are equal")
	}

	var (
		xa   int
		xb   int
		st   XY
		line string
	)

	if pos1.X <= pos2.X && pos1.Y <= pos2.Y {
		xa = pos2.X - pos1.X
		xb = pos2.Y - pos1.Y
		st = pos1
	} else if pos1.X >= pos2.X && pos1.Y >= pos2.Y {
		xa = pos1.X - pos2.X
		xb = pos1.Y - pos2.Y
		st = pos2
	} else if pos1.X >= pos2.X && pos1.Y <= pos2.Y {
		xa = pos1.X - pos2.X
		xb = pos2.Y - pos1.Y
		st = XY{
			X: pos2.X,
			Y: pos1.Y,
		}
	} else if pos1.X <= pos2.X && pos1.Y >= pos2.Y {
		xa = pos2.X - pos1.X
		xb = pos1.Y - pos2.Y
		st = XY{
			X: pos1.X,
			Y: pos2.Y,
		}
	}

	if xa == 0 {
		xa = 1
	}
	line = strings.Repeat(string(ch), xa)

	CursorTo(st)
	for i := 0; i < xb; i++ {
		fmt.Print(line)
		CursorDown(1)
		CursorLeft(xa)
	}

	return nil
}

func ColorArea256(c uint8, pos1 XY, pos2 XY) error {

	Color256Bg(c)
	defer SetGRA(BgDefault, FgDefault)

	err := CharArea(' ', pos1, pos2)
	if err != nil {
		return err
	}

	return nil
}

func ColorAreaRGB(c RGB, pos1 XY, pos2 XY) error {
	ColorRGBBg(c)
	defer SetGRA(BgDefault, FgDefault)

	err := CharArea(' ', pos1, pos2)
	if err != nil {
		return err
	}

	return nil
}
