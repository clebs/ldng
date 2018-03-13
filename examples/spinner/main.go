package main

import (
	"time"

	"github.com/clebs/ldng"
	"github.com/clebs/ldng/term"
)

func main() {
	s := ldng.NewSpin(ldng.SpinPrefix("Processing"), ldng.SpinSuccess("\nSuccess!!\n"))

	term.HideCursor()
	s.Start()
	time.Sleep(time.Second * 8)
	s.Stop()
	term.ShowCursor()
}
