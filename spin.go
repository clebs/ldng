package ldng

import (
	"fmt"
	"time"

	"github.com/clebs/ldng/term"
)

// Spin represents a Spin bar in the term
type Spin struct {
	frames  []string
	period  time.Duration
	prefix  string
	success string
	stop    chan struct{}
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
		},
		period: time.Second,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Start the spinner
func (s *Spin) Start() {
	s.stop = make(chan struct{}, 0)
	go func() {
		frame := 0
		for {
			select {
			case <-time.After(s.period):
				term.Clearln()
				if s.prefix != "" {
					fmt.Printf("%s ", s.prefix)
				}
				fmt.Printf("%s", s.frames[frame])
				frame = (frame + 1) % len(s.frames)
			case <-s.stop:
				return
			}
		}
	}()
}

// Stop the spinner
func (s *Spin) Stop() {
	close(s.stop)
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
