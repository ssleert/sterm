package sterm

import (
	"fmt"
	"github.com/ssleert/sfolib"
)

// a simple function for easy formatting of the escape sequence
func escape(f string, args ...interface{}) string {
	return fmt.Sprintf(
		"%s%s",
		Esc,
		fmt.Sprintf(f, args...),
	)
}

// get integer len in characters
func chLen(a int) int {
	switch {
	case a > 999999:
		return 7
	case a > 99999:
		return 6
	case a > 9999:
		return 5
	case a > 999:
		return 4
	case a > 99:
		return 3
	case a > 9:
		return 2
	default:
		return 1
	}
}

// find the length and width of a rectangle by 2 points
func findLH(p1x, p1y, p2x, p2y int) (int, int, int, int, error) {
	if p1x+p1y == p2x+p2y {
		return 0, 0, 0, 0, ErrPositionsEqual
	}

	var (
		xa int
		xb int
		st [2]int
	)

	xa = sfolib.Abs(p1x-p2x) + 1
	xb = sfolib.Abs(p1y-p2y) + 1

	if p1x <= p2x && p1y <= p2y {
		st = [2]int{p1x, p1y}
	} else if p1x >= p2x && p1y >= p2y {
		st = [2]int{p2x, p2y}
	} else if p1x >= p2x && p1y <= p2y {
		st = [2]int{p2x, p1y}
	} else if p1x <= p2x && p1y >= p2y {
		st = [2]int{p1x, p2y}
	}

	return xa, xb, st[0], st[1], nil
}
