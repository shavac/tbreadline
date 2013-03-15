package tbreadline

import (
	termbox "github.com/nsf/termbox-go"
)

type editline struct {
	Prompt func() *string

	LineBuf string
	// position beginning after immediate prompt
	pos int
}

func (l *editline) String() *string {
	return &l.LineBuf
}

func (l *editline) len() int {
	return len(l.LineBuf)
}

func newLine(Prompt func()(*string)) *editline {
	return &editline {
		Prompt: Prompt,
		LineBuf: "",
		pos: 0,
	}
}

func PushKey(k termbox.Key) {
	
}