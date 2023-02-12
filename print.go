package sterm

import (
	"github.com/ssleert/sfolib"
	"strings"
)

// coz normal print prints from left to right
// RevPrint return string to print from right to left
func RevPrint(v string) string {
	var (
		s strings.Builder
		l = len(v)
	)
	s.Grow(l + 4 + chLen(l-1))

	if l > 1 {
		s.WriteString(CursorLeft(l - 1))
	} else {
		s.WriteString(CursorLeft(2))
	}
	s.WriteString(v)
	return s.String()
}

// print characters between 2 points
func CharArea(ch rune, p1x, p1y, p2x, p2y int) (string, error) {
	xa, xb, stx, sty, err := findLH(p1x, p1y, p2x, p2y)
	if err != nil {
		return "", err
	}

	var s strings.Builder
	s.Grow(xa*xb + xb*(4+chLen(xa)))

	line := strings.Repeat(string(ch), xa)

	s.WriteString(CursorTo(stx, sty))
	for i := 0; i < xb; i++ {
		s.WriteString(line)
		s.WriteString(CursorDown(1))
		s.WriteString(CursorLeft(xa))
	}

	return s.String(), nil
}

// frame the area between 2 points
func FrameArea(sym Borders, p1x, p1y, p2x, p2y int) (string, error) {
	xa, xb, stx, sty, err := findLH(p1x, p1y, p2x, p2y)
	if err != nil {
		return "", err
	}

	if xa == 1 || xb == 1 {
		return "", ErrLWOfAreaUnderTwo
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
	lineBuf.Reset()

	var s strings.Builder
	s.Grow(xb + xb*((4+chLen(xa-2))+(4+chLen(xa))) + xb*5 + xb*2 + chLen(
		stx) + chLen(sty) + 5 + len(lineTop) + len(lineBot))

	s.WriteString(CursorTo(stx, sty))
	s.WriteString(lineTop)
	s.WriteString(CursorLeft(xa))
	s.WriteString(CursorDown(1))
	for i := 0; i < xb-2; i++ {
		s.WriteRune(sym[0])
		s.WriteString(CursorRight(xa - 2))
		s.WriteRune(sym[0])
		s.WriteString(CursorDown(1))
		s.WriteString(CursorLeft(xa))
	}
	s.WriteString(lineBot)

	return s.String(), nil
}

func ReserveArea(n int) string {
	var s strings.Builder
	n = sfolib.Abs(n)
	s.Grow(n + (4 + chLen(n)))
	s.WriteString(strings.Repeat("\n", n))
	s.WriteString(CursorUp(n))
	return s.String()
}

func DrawTable(tbl [][]string, sym Borders) ([]string, error) {
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
	var s strings.Builder
	s.Grow(lengthsSum * 4)

	s.WriteRune(sym[2])
	for i := 0; i < lengthsSum; i++ {
		s.WriteRune(sym[1])
	}
	s.WriteRune(sym[3])

	lines = append(lines, s.String())
	s.Reset()
	s.Grow(lengthsSum * 4)
	for r, line := range tbl {
		for i, e := range line {
			if i == 0 {
				s.WriteRune(sym[0])
			}
			for p := 0; p < lengths[i]-len(e)+1; p++ {
				s.WriteRune(' ')
			}
			s.WriteString(e)
			s.WriteRune(' ')
			s.WriteRune(sym[0])
		}
		lines = append(lines, s.String())
		if r == 0 {
			s.Reset()
			s.Grow(lengthsSum * 4)
			s.WriteRune(sym[0])
			for i := 0; i < lengthsSum; i++ {
				s.WriteRune(sym[1])
			}
			s.WriteRune(sym[0])
			lines = append(lines, s.String())
		}
		s.Reset()
		s.Grow(lengthsSum * 4)
	}
	s.WriteRune(sym[4])
	for i := 0; i < lengthsSum; i++ {
		s.WriteRune(sym[1])
	}
	s.WriteRune(sym[5])
	lines = append(lines, s.String())

	return lines, nil
}
