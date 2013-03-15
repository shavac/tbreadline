package tbreadline

import (
	termbox "github.com/nsf/termbox-go"
)

type keymap map[termbox.Key]func(*terminal)

var keymap_default = keymap{
	termbox.KeyCtrlL: rl_clear_screen,
	termbox.KeyBackspace: rl_backspace,
}

var Keymap = keymap_default

func SetKeyMap(km keymap) {
	Keymap = km
}
