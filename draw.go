package sterm

import (
	"fmt"
	"strings"
)

// print characters between 2 points
func CharArea(ch rune, pos1 XY, pos2 XY) error {
	xa, xb, st, err := findLH(pos1, pos2)
	if err != nil {
		return err
	}

	if xa == 0 {
		xa = 1
	}
	line := strings.Repeat(string(ch), xa)

	CursorTo(st)
	for i := 0; i < xb; i++ {
		fmt.Print(line)
		CursorDown(1)
		CursorLeft(xa)
	}

	return nil
}

// frame the area between 2 points
func FrameArea(sym [6]rune, pos1, pos2 XY) error {
	xa, xb, st, err := findLH(pos1, pos2)
	if err != nil {
		return err
	}

	if xa == 1 || xb == 1 {
		return ErrLWOfAreaUnderTwo
	}

	lineTop := string(sym[2]) + strings.Repeat(string(sym[1]), xa-2) + string(sym[3])
	lineBot := string(sym[4]) + strings.Repeat(string(sym[1]), xa-2) + string(sym[5])

	CursorTo(st)
	fmt.Print(lineTop)
	CursorLeft(xa)
	CursorDown(1)
	for i := 0; i < xb-2; i++ {
		fmt.Print(string(sym[0]))
		CursorRight(xa - 2)
		fmt.Print(string(sym[0]))
		CursorDown(1)
		CursorLeft(xa)
	}
	fmt.Print(lineBot)

	return nil
}

func ReserveArea(n int) error {
	if n < 0 {
		return ErrNegative
	}

	fmt.Print(strings.Repeat("\n", n))
	CursorUp(n)

	return nil
}
