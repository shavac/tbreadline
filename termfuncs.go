package tbreadline

import (
//termbox "github.com/nsf/termbox-go"
)

func rl_clear_screen(t *terminal) {
}

func rl_backspace(t *terminal) {
	if t.line.len() <= 0 {
		rl_ding(t)
	} else {
		t.line.LineBuf = t.line.LineBuf[:t.line.pos-1] + t.line.LineBuf[t.line.pos:]
		t.line.pos -= 1
		t.Flush()
	}
}

func rl_ding(t *terminal) {
	t.c.Write([]byte("\007"))
}
