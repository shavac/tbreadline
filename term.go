package tbreadline

import (
	"sync"
	"io"

	termbox "github.com/nsf/termbox-go"
)

// Terminal contains the state for running a VT100 terminal that is capable of
// reading lines of input.
type terminal struct {
	// AutoCompleteCallback, if non-null, is called for each keypress
	// with the full input line and the current position of the cursor.
	// If it returns a nil newLine, the key press is processed normally.
	// Otherwise it returns a replacement line and the new cursor position.
	AutoCompleteCallback func(line editline, key termbox.Key) (newLine editline)

	// Prompt, returns command line prompt
	Prompt func() (prompt *string)

	// Escape contains a pointer to the escape codes for this terminal.
	// It's always a valid pointer, although the escape codes themselves
	// may be empty if the terminal doesn't support them.
	//Escape *EscapeCodes

	// lock protects the terminal and the state in this object from
	// concurrent processing of a key press and a Write() call.
	lock sync.Mutex

	c io.ReadWriter

	// line is the current line being entered.
	line *editline

	// echo is true if local echo is enabled
	echo bool

	// cursorX contains the current X value of the cursor where the left
	// edge is 0. cursorY contains the row number where the first row of
	// the current line is 0.
	cursorX, cursorY int

}

func (t *terminal) init() {
	termbox.Init()
}

// Finalizes tbreadline library, should be called after successful initialization
// when termbox's functionality isn't required anymore.
func (t *terminal) close() {
	termbox.Close()
}

func (t *terminal) reset() error {
	return termbox.Init()
}

func (t *terminal) Size() (termWidth, termHeight int) {
	return termbox.Size()
}

func (t *terminal) NewLine() *editline {
	return newLine(t.Prompt)
}

func (t *terminal) Flush() {
}