package sterm

import (
	"errors"
	"fmt"
)

// a simple function for easy formatting and printing of the escape sequence
func escape(f string, args ...any) {
	fmt.Printf(
		"%s%s",
		Esc,
		fmt.Sprintf(f, args...),
	)
}

// find the length and width of a rectangle by 2 points
func findLH(pos1, pos2 XY) (int, int, XY, error) {
	if pos1 == pos2 {
		return 0, 0, XY{}, errors.New("the positions of the points are equal")
	}

	var (
		xa uint
		xb uint
		st XY
	)

	if pos1.X <= pos2.X && pos1.Y <= pos2.Y {
		xa = pos2.X - pos1.X + 1
		xb = pos2.Y - pos1.Y + 1
		st = pos1
	} else if pos1.X >= pos2.X && pos1.Y >= pos2.Y {
		xa = pos1.X - pos2.X + 1
		xb = pos1.Y - pos2.Y + 1
		st = pos2
	} else if pos1.X >= pos2.X && pos1.Y <= pos2.Y {
		xa = pos1.X - pos2.X + 1
		xb = pos2.Y - pos1.Y + 1
		st = XY{
			X: pos2.X,
			Y: pos1.Y,
		}
	} else if pos1.X <= pos2.X && pos1.Y >= pos2.Y {
		xa = pos2.X - pos1.X + 1
		xb = pos1.Y - pos2.Y + 1
		st = XY{
			X: pos1.X,
			Y: pos2.Y,
		}
	}

	return int(xa), int(xb), st, nil
}
