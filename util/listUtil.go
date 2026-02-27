package util

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Updates the list given the list and tea msg recived.
func ListUpdate(list list.Model, msg tea.Msg) (list.Model, tea.Cmd) {
	if len(list.Items()) == 0 {
		return list, nil
	}
	return list.Update(msg)
}

// Sets desired defaults of a list like the title and infinate scrolling.
func ListDefaults(list list.Model) list.Model {
	list.Title = "Desired Operation:"
	list.InfiniteScrolling = true
	return list
}
