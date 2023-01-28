package sterm

import "errors"

var (
	ErrPositionsEqual        = errors.New("the positions of the points are equal")
	ErrNegative              = errors.New("value is negative")
	ErrLWOfAreaUnderTwo      = errors.New("the length and width of the area must be at least 2")
	ErrTblLineShorterThanTbl = errors.New("the table line is shorter than the table")
)
