package sterm

import "fmt"

// set graphics rendition attributes for terminal
func SetGRA(sgras ...string) {
	for _, s := range sgras {
		fmt.Print(s)
	}
}
