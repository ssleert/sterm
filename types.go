package sterm

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
	Reset      = Esc + "[1;0m"
	Bold       = Esc + "[1;1m"
	Faint      = Esc + "[1;2m"
	Italic     = Esc + "[1;3m"
	Underlined = Esc + "[1;4m"
	BlinkSlow  = Esc + "[1;5m"
	BlinkRapid = Esc + "[1;6m"
	Reverse    = Esc + "[1;7m"
	Invisible  = Esc + "[1;8m"
	Strike     = Esc + "[1;9m"
	Overline   = Esc + "[1;53m"

	// foreground colors
	FgDefault      = Esc + "[1;39m"
	FgBlack        = Esc + "[1;30m"
	FgRed          = Esc + "[1;31m"
	FgGreen        = Esc + "[1;32m"
	FgYellow       = Esc + "[1;33m"
	FgBlue         = Esc + "[1;34m"
	FgMagenta      = Esc + "[1;35m"
	FgCyan         = Esc + "[1;36m"
	FgLightGray    = Esc + "[1;37m"
	FgDarkGray     = Esc + "[1;90m"
	FgLightRed     = Esc + "[1;91m"
	FgLightGreen   = Esc + "[1;92m"
	FgLightYellow  = Esc + "[1;93m"
	FgLightBlue    = Esc + "[1;94m"
	FgLightMagenta = Esc + "[1;95m"
	FgLightCyan    = Esc + "[1;96m"
	FgWhite        = Esc + "[1;97m"

	// background colors
	BgDefault      = Esc + "[1;49m"
	BgBlack        = Esc + "[1;40m"
	BgRed          = Esc + "[1;41m"
	BgGreen        = Esc + "[1;42m"
	BgYellow       = Esc + "[1;43m"
	BgBlue         = Esc + "[1;44m"
	BgMagenta      = Esc + "[1;45m"
	BgCyan         = Esc + "[1;46m"
	BgLightGray    = Esc + "[1;47m"
	BgDarkGray     = Esc + "[1;100m"
	BgLightRed     = Esc + "[1;101m"
	BgLightGreen   = Esc + "[1;102m"
	BgLightYellow  = Esc + "[1;103m"
	BgLightBlue    = Esc + "[1;104m"
	BgLightMagenta = Esc + "[1;105m"
	BgLightCyan    = Esc + "[1;106m"
	BgWhite        = Esc + "[1;107m"
)
