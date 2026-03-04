package util

import (
	"BCLAC/enums"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
)

func SetInputTextDefaults(ti *textinput.Model) {
	ti.Placeholder = enums.TEXT_PLACEHOLDER
	ti.Focus()
	ti.CharLimit = enums.TEXT_CHAR_LIMT
	ti.Width = enums.TEXT_CHAR_LIMT
	ti.Cursor.SetMode(cursor.CursorBlink)
}
