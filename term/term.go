// Package term provides functionality to manipulate the terminal
package term

import "fmt"

//HideCursor hides the cursor
func HideCursor() {
	fmt.Print("\033[?25l")
}

// ShowCursor displays the cursor if it was hidden
func ShowCursor() {
	fmt.Print("\033[?25h")
}

// Clearln moves the cursor to the start of the line and deletes the whole line
func Clearln() {
	fmt.Print("\r\033[K")
}
