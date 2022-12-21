package sterm

import (
	"fmt"
)

// a simple function for easy formatting of the escape sequence
func escape(f string, args ...any) {
	fmt.Printf(
		"%s%s",
		Esc,
		fmt.Sprintf(f, args...),
	)
}
