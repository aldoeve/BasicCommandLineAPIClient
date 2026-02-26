package util

import (
	"github.com/charmbracelet/bubbles/list"

	"BCLAC/styles"
)

func HandleResize(list list.Model) list.Model {
	if len(list.Items()) == 0 {
		return list
	}
	x, y := styles.DocStyle.GetFrameSize()

	const (
		height = 20
		width  = 20
	)

	list.SetSize(width-x, height-y)
	return list
}
