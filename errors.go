package sterm

import "errors"

var (
	ErrColorOutOfRange  = errors.New("color value is negative or greater than 255")
	ErrPositionsEqual   = errors.New("the positions of the points are equal")
	ErrNegative         = errors.New("value is negative")
	ErrLWOfAreaUnderTwo = errors.New("the length and width of the area must be at least 2")
)
