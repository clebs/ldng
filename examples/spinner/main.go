package main

import (
	"time"

	"github.com/clebs/ldng"
	"github.com/clebs/ldng/term"
)

func main() {
	s := ldng.NewSpin(ldng.SpinPrefix("Processing"), ldng.SpinPeriod(100*time.Millisecond), ldng.SpinSuccess("Success!!\n"))

	term.HideCursor()
	stop := s.Start()
	time.Sleep(time.Second * 3)
	stop <- struct{}{} // stop the spinner after 3 seconds
	<-stop             // wait for the spinner to finish the stop task
	term.ShowCursor()
}
