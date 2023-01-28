package sterm

import (
	"fmt"
	"math"
)

// a simple function for easy formatting of the escape sequence
func escape(f string, args ...any) string {
	return fmt.Sprintf(
		"%s%s",
		Esc,
		fmt.Sprintf(f, args...),
	)
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

	xa = int(math.Abs(float64(p1x - p2x))) + 1
	xb = int(math.Abs(float64(p1y - p2y))) + 1

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
