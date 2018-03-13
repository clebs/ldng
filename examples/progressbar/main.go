package main

import (
	"time"

	"github.com/clebs/ldng"
	"github.com/clebs/ldng/term"
)

func main() {
	p := ldng.NewProgress(ldng.ProgressPrefix("Processing"), ldng.ProgressSuccess("\nSuccess!!\n"))

	term.HideCursor()
	for i := 0; i <= 10; i++ {
		p.Update(i * 10)
		time.Sleep(time.Second)
	}
	term.ShowCursor()
}
