package util

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func ListUpdate(list list.Model, msg tea.Msg) (list.Model, tea.Cmd) {
	if len(list.Items()) == 0 {
		return list, nil
	}
	return list.Update(msg)
}
