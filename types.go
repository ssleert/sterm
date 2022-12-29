package sterm

type XY struct {
	X uint
	Y uint
}

type RGB struct {
	R uint8
	G uint8
	B uint8
}

// chars for borders
var (
	RoundedBorders = [6]rune{'│', '─', '╭', '╮', '╰', '╯'}
	SharpBorders   = [6]rune{'│', '─', '┌', '┐', '└', '┘'}
	DoubleBorders  = [6]rune{'║', '═', '╔', '╗', '╚', '╝'}
	AsciiBorders   = [6]rune{'|', '-', '+', '+', '+', '+'}
	DotBorders     = [6]rune{'.', '.', '.', '.', '.', '.'}
)

// escape
const Esc = "\x1b"

const (
	// styles
	Reset      = "\x1b[1;0m"
	Bold       = "\x1b[1;1m"
	Faint      = "\x1b[1;2m"
	Italic     = "\x1b[1;3m"
	Underlined = "\x1b[1;4m"
	BlinkSlow  = "\x1b[1;5m"
	BlinkRapid = "\x1b[1;6m"
	Reverse    = "\x1b[1;7m"
	Invisible  = "\x1b[1;8m"
	Strike     = "\x1b[1;9m"
	Overline   = "\x1b[1;53m"

	// foreground colors
	FgDefault      = "\x1b[1;39m"
	FgBlack        = "\x1b[1;30m"
	FgRed          = "\x1b[1;31m"
	FgGreen        = "\x1b[1;32m"
	FgYellow       = "\x1b[1;33m"
	FgBlue         = "\x1b[1;34m"
	FgMagenta      = "\x1b[1;35m"
	FgCyan         = "\x1b[1;36m"
	FgLightGray    = "\x1b[1;37m"
	FgDarkGray     = "\x1b[1;90m"
	FgLightRed     = "\x1b[1;91m"
	FgLightGreen   = "\x1b[1;92m"
	FgLightYellow  = "\x1b[1;93m"
	FgLightBlue    = "\x1b[1;94m"
	FgLightMagenta = "\x1b[1;95m"
	FgLightCyan    = "\x1b[1;96m"
	FgWhite        = "\x1b[1;97m"

	// background colors
	BgDefault      = "\x1b[1;49m"
	BgBlack        = "\x1b[1;40m"
	BgRed          = "\x1b[1;41m"
	BgGreen        = "\x1b[1;42m"
	BgYellow       = "\x1b[1;43m"
	BgBlue         = "\x1b[1;44m"
	BgMagenta      = "\x1b[1;45m"
	BgCyan         = "\x1b[1;46m"
	BgLightGray    = "\x1b[1;47m"
	BgDarkGray     = "\x1b[1;100m"
	BgLightRed     = "\x1b[1;101m"
	BgLightGreen   = "\x1b[1;102m"
	BgLightYellow  = "\x1b[1;103m"
	BgLightBlue    = "\x1b[1;104m"
	BgLightMagenta = "\x1b[1;105m"
	BgLightCyan    = "\x1b[1;106m"
	BgWhite        = "\x1b[1;107m"
)
