package views

import (
	"github.com/charmbracelet/bubbles/list"

	"BCLAC/enums"
	"BCLAC/util"
)

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

const (
	MAINSCREEN_OpTIONS_COUNT = 4
)

// Returns the first list of the application.
func InitiateFirstSelectionList(isAltscreen bool) list.Model {
	items := []list.Item{
		item{title: "Paste & Fire"},
		item{title: "Quick Build"},
		item{title: "Display Recents"},
		item{title: "Manual SQL"},
	}
	choices := list.New(items, list.NewDefaultDelegate(), 0, 0)
	choices = util.HandleResizeofList(choices, isAltscreen)
	choices = util.SetListDefaults(choices)

	return choices
}

func ReturnNextStateFromMainSelection(index uint) enums.State {
	options := [MAINSCREEN_OpTIONS_COUNT]enums.State{
		enums.PASTE,
		enums.QUICK_BUILD,
		enums.RECENTS,
		enums.MANUAL,
	}
	if index >= MAINSCREEN_OpTIONS_COUNT {
		index = 0
	}
	return options[index]
}

// Returns first view after the intro screen.
func FirstSelectionView(list list.Model) string {
	return list.View()
}
