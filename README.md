# LoadinGo

LoadinGo, abbreviated `ldng`, is a package that allows to easily create loading indicators for the terminal in Go programs.

## Examples

### Progress Bar
```go
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

```

### Spinner
```go
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

```