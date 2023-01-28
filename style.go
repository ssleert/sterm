package sterm

// set graphics rendition attributes for terminal
func SetGRA(gras ...string) string {
	if len(gras) > 1 {
		return ""
	}

	var strGras string
	for _, s := range gras {
		strGras += s
	}
	return strGras
}

func Color256Fg(c int) string       { return escape("[38;5;%dm", c) }
func Color256Bg(c int) string       { return escape("[48;5;%dm", c) }
func ColorRGBFg(r, g, b int) string { return escape("[38;2;%d;%d;%dm", r, g, b) }
func ColorRGBBg(r, g, b int) string { return escape("[48;2;%d;%d;%dm", r, g, b) }
