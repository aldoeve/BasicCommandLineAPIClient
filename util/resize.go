package util

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/x/term"

	"BCLAC/enums"
	"BCLAC/styles"
)

func HandleResizeofList(list list.Model, isAltscreen bool) list.Model {
	if len(list.Items()) == 0 {
		return list
	}
	x, y := styles.DocStyle.GetFrameSize()
	width, height := GetTerminalSize()

	if isAltscreen == false {
		height = min(enums.DEFAULT_HEIGHT, height)
	}

	list.SetSize(width-x, height-y)
	return list
}

func GetTerminalSize() (int, int) {
	width, height, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		width = enums.DEFAULT_WIDTH
		height = enums.DEFAULT_HEIGHT
	}
	return width, height
}
