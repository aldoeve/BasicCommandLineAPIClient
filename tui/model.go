// Package tui holds the main logic to the application.
package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"BCLAC/enums"
	"BCLAC/util"
)

type Model struct {
	state      enums.State
	altscreen  bool
	suspending bool
	list       list.Model
	textInput  textinput.Model
	spinner    spinner.Model
}

func InitalModel() *Model {
	ti := textinput.New()
	return &Model{textInput: util.SetInputTextDefaults(ti)}
}

func (m Model) Init() tea.Cmd {
	return nil
}
