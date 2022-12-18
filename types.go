package sterm

// escape
const Esc = "\x1b"

const (
	// styles
	Reset      = "\033[1;0m"
	Bold       = "\033[1;1m"
	Faint      = "\033[1;2m"
	Italic     = "\033[1;3m"
	Underlined = "\033[1;4m"
	BlinkSlow  = "\033[1;5m"
	BlinkRapid = "\033[1;6m"
	Reverse    = "\033[1;7m"
	Invisible  = "\033[1;8m"
	Strike     = "\033[1;9m"
	Overline   = "\033[1;53m"

	// foreground colors
	FgDefault      = "\033[1;39m"
	FgBlack        = "\033[1;30m"
	FgRed          = "\033[1;31m"
	FgGreen        = "\033[1;32m"
	FgYellow       = "\033[1;33m"
	FgBlue         = "\033[1;34m"
	FgMagenta      = "\033[1;35m"
	FgCyan         = "\033[1;36m"
	FgLightGray    = "\033[1;37m"
	FgDarkGray     = "\033[1;90m"
	FgLightRed     = "\033[1;91m"
	FgLightGreen   = "\033[1;92m"
	FgLightYellow  = "\033[1;93m"
	FgLightBlue    = "\033[1;94m"
	FgLightMagenta = "\033[1;95m"
	FgLightCyan    = "\033[1;96m"
	FgWhite        = "\033[1;97m"

	// background colors
	BgDefault      = "\033[1;49m"
	BgBlack        = "\033[1;40m"
	BgRed          = "\033[1;41m"
	BgGreen        = "\033[1;42m"
	BgYellow       = "\033[1;43m"
	BgBlue         = "\033[1;44m"
	BgMagenta      = "\033[1;45m"
	BgCyan         = "\033[1;46m"
	BgLightGray    = "\033[1;47m"
	BgDarkGray     = "\033[1;100m"
	BgLightRed     = "\033[1;101m"
	BgLightGreen   = "\033[1;102m"
	BgLightYellow  = "\033[1;103m"
	BgLightBlue    = "\033[1;104m"
	BgLightMagenta = "\033[1;105m"
	BgLightCyan    = "\033[1;106m"
	BgWhite        = "\033[1;107m"
)
