package ldng

import (
	"bytes"
	"fmt"
	"time"

	"github.com/clebs/ldng/term"
)

// Spin represents a Spin bar in the term
type Spin struct {
	frames  []string
	current int
	period  time.Duration
	prefix  string
	success string
}

// NewSpin creates a new Spin
func NewSpin(opts ...func(*Spin)) *Spin {
	// defaults
	s := &Spin{
		frames: []string{
			"|",
			"/",
			"-",
			"\\",
			"|",
			"/",
			"-",
			"\\",
		},
		period: time.Second,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Spin) frame() string {
	s.current = (s.current + 1) % len(s.frames)
	return s.frames[s.current]
}

func (s *Spin) String() string {
	var b bytes.Buffer
	if s.prefix != "" {
		b.WriteString(fmt.Sprintf("%s ", s.prefix))
	}
	b.WriteString(fmt.Sprintf("%s", s.frame()))

	return b.String()
}

// Start the spinner and return a channel to stop it.
// Use the stop channel to stop the spinner.
// Once it is stopped, the spinner will close the channel, so it can be used to wait for the spinner to finish.
func (s *Spin) Start() (stop chan struct{}) {
	stop = make(chan struct{}, 0)
	go func() {
		for {
			select {
			case <-time.After(s.period):
				term.Clearln()
				fmt.Print(s.String())
			case <-stop:
				if s.success != "" {
					term.Clearln()
					fmt.Print(s.success)
				}
				close(stop) // callers can receive on this channel to wait for the spinner to stop
				return
			}
		}
	}()
	return
}

/* Options */

// SpinPrefix sets the prefix to be displayed before a Spinner
func SpinPrefix(pre string) func(*Spin) {
	return func(s *Spin) {
		s.prefix = pre
	}
}

// SpinSuccess sets the success message to be displayed when the spinner is stopped.
func SpinSuccess(sc string) func(*Spin) {
	return func(s *Spin) {
		s.success = sc
	}
}

// SpinFrames sets the frames forming the Spin bar.
func SpinFrames(f []string) func(*Spin) {
	return func(p *Spin) {
		p.frames = f
	}
}

// SpinPeriod sets the speed at which the Spinner  spins.
func SpinPeriod(p time.Duration) func(*Spin) {
	return func(s *Spin) {
		s.period = p
	}
}
