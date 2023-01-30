package sterm

import (
	"fmt"
	"strings"
)

// coz normal print prints from left to right
// RevPrint prints from right to left
func RevPrint(a ...any) {
	s := fmt.Sprint(a...)
	fmt.Print(CursorLeft(len(s) - 1))
	fmt.Print(s)
}

// print characters between 2 points
func CharArea(ch rune, p1x, p1y, p2x, p2y int) error {
	xa, xb, stx, sty, err := findLH(p1x, p1y, p2x, p2y)
	if err != nil {
		return err
	}

	line := strings.Repeat(string(ch), xa)

	fmt.Print(CursorTo(stx, sty))
	for i := 0; i < xb; i++ {
		fmt.Print(line)
		fmt.Print(CursorDown(1))
		fmt.Print(CursorLeft(xa))
	}

	return nil
}

// frame the area between 2 points
func FrameArea(sym [6]rune, p1x, p1y, p2x, p2y int) error {
	xa, xb, stx, sty, err := findLH(p1x, p1y, p2x, p2y)
	if err != nil {
		return err
	}

	if xa == 1 || xb == 1 {
		return ErrLWOfAreaUnderTwo
	}

	var lineBuf strings.Builder
	lineBuf.Grow(xa)

	lineBuf.WriteRune(sym[2])
	for i := 0; i < xa-2; i++ {
		lineBuf.WriteRune(sym[1])
	}
	lineBuf.WriteRune(sym[3])

	lineTop := lineBuf.String()

	lineBuf.Reset()
	lineBuf.Grow(xa)
	lineBuf.WriteRune(sym[4])
	for i := 0; i < xa-2; i++ {
		lineBuf.WriteRune(sym[1])
	}
	lineBuf.WriteRune(sym[5])

	lineBot := lineBuf.String()

	fmt.Print(CursorTo(stx, sty))
	fmt.Print(lineTop)
	fmt.Print(CursorLeft(xa))
	fmt.Print(CursorDown(1))
	for i := 0; i < xb-2; i++ {
		fmt.Print(string(sym[0]))
		fmt.Print(CursorRight(xa - 2))
		fmt.Print(string(sym[0]))
		fmt.Print(CursorDown(1))
		fmt.Print(CursorLeft(xa))
	}
	fmt.Print(lineBot)

	return nil
}

func ReserveArea(n int) error {
	if n < 0 {
		return ErrNegative
	}

	fmt.Print(strings.Repeat("\n", n))
	fmt.Println(CursorUp(n))
	return nil
}

func DrawTable(tbl [][]string, sym [6]rune) ([]string, error) {
	for i, e := range tbl {
		if len(e) != len(tbl[len(tbl)-i-1]) {
			return nil, ErrTblLineShorterThanTbl
		}
	}

	lengths := make([]int, 0, len(tbl[0]))
	var length int
	for i := range tbl[0] {
		for _, e := range tbl {
			if len(e[i]) > length {
				length = len(e[i])
			}
		}
		lengths = append(lengths, length)
		length = 0
	}

	lengthsSum := -1
	for _, e := range lengths {
		lengthsSum += e + 3
	}

	lines := make([]string, 0, len(tbl)+3)
	var lineStr strings.Builder
	lineStr.Grow(lengthsSum*3 + 30)

	lineStr.WriteRune(sym[2])
	for i := 0; i < lengthsSum; i++ {
		lineStr.WriteRune(sym[1])
	}
	lineStr.WriteRune(sym[3])

	lines = append(lines, lineStr.String())
	lineStr.Reset()
	lineStr.Grow(lengthsSum*3 + 30)
	for r, line := range tbl {
		for i, e := range line {
			if i == 0 {
				lineStr.WriteRune(sym[0])
			}
			for p := 0; p < lengths[i]-len(e)+1; p++ {
				lineStr.WriteRune(' ')
			}
			lineStr.WriteString(e)
			lineStr.WriteRune(' ')
			lineStr.WriteRune(sym[0])
		}
		lines = append(lines, lineStr.String())
		if r == 0 {
			lineStr.Reset()
			lineStr.Grow(lengthsSum*3 + 30)
			lineStr.WriteRune(sym[0])
			for i := 0; i < lengthsSum; i++ {
				lineStr.WriteRune(sym[1])
			}
			lineStr.WriteRune(sym[0])
			lines = append(lines, lineStr.String())
		}
		lineStr.Reset()
		lineStr.Grow(lengthsSum*3 + 30)
	}
	lineStr.WriteRune(sym[4])
	for i := 0; i < lengthsSum; i++ {
		lineStr.WriteRune(sym[1])
	}
	lineStr.WriteRune(sym[5])
	lines = append(lines, lineStr.String())

	return lines, nil
}
