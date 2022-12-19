package sterm

// clear the line to the right of the cursor
func ClearLineRight() { escape("[0K") }

// clear the line to the left of the cursor
func ClearLineLeft() { escape("[1K") }

// clear the entire line under the cursor
func ClearEntireLine() { escape("[2K") }

// clear everything below the cursor
func ClearScreenDown() { escape("[0J") }

// clear everything above the cursor
func ClearScreenUp() { escape("[1J") }

// clear everything on the screen
func ClearEntireScreen() { escape("[2J") }

func ClearHome() {
	ClearEntireScreen()
	CursorHome()
}
