package ldng

import (
	"bytes"
	"fmt"

	"github.com/clebs/ldng/term"
)

// Progress represents a progress bar in the term
type Progress struct {
	frames     []string
	completion int
	prefix     string
	success    string
}

// NewProgress creates a new Progress
func NewProgress(opts ...func(*Progress)) *Progress {
	// defaults
	p := &Progress{
		frames: []string{
			"[          ]",
			"[=         ]",
			"[==        ]",
			"[===       ]",
			"[====      ]",
			"[=====     ]",
			"[======    ]",
			"[=======   ]",
			"[========  ]",
			"[========= ]",
			"[==========]",
		},
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *Progress) frame(percent int) string {
	normalized := percent * (len(p.frames) - 1) / 100
	return p.frames[normalized]
}

// String returns a string representation of the progress bar in its current state
func (p *Progress) String() string {
	var b bytes.Buffer

	if p.prefix != "" {
		b.WriteString(fmt.Sprintf("%s ", p.prefix))
	}

	b.WriteString(fmt.Sprintf("%s", p.frame(p.completion)))

	if p.completion == 100 && p.success != "" {
		b.WriteString(fmt.Sprintf("%s", p.success))
	}
	return b.String()
}

// Update the completion percentage of the progress, triggering a rendition
func (p *Progress) Update(percent int) {
	p.completion = percent
	term.Clearln()
	fmt.Print(p.String())
}

/* Options */

// ProgressPrefix sets the prefix to be displayed before a progress bar
func ProgressPrefix(pre string) func(*Progress) {
	return func(p *Progress) {
		p.prefix = pre
	}
}

// ProgressSuccess sets the success message to be displayed when finished successfully.
func ProgressSuccess(sc string) func(*Progress) {
	return func(p *Progress) {
		p.success = sc
	}
}

// ProgressFrames sets the frames forming the progress bar.
func ProgressFrames(f []string) func(*Progress) {
	return func(p *Progress) {
		p.frames = f
	}
}
