package views

import (
	"github.com/charmbracelet/bubbles/list"

	"BCLAC/util"
)

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

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
	choices = util.ListDefaults(choices)

	return choices
}

// Returns first view after the intro screen.
func FirstSelectionView(list list.Model) string {
	return list.View()
}
