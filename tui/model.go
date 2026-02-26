package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"BCLAC/enums"
)

type Model struct {
	state      enums.State
	altscreen  bool
	suspending bool
	list       list.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}
