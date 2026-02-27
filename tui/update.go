package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"BCLAC/enums"
	"BCLAC/tui/views"
	"BCLAC/util"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.ResumeMsg:
		m.suspending = false
		return m, nil
	case tea.WindowSizeMsg:
		m.list = util.HandleResizeofList(m.list, m.altscreen)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "right":
			m.state = enums.MAIN_VIEW
			//cmd := clearScreenInAltMode(m.altscreen)
			m.list = views.InitiateFirstSelectionList(m.altscreen)
			return m, nil //cmd
		case "ctrl+q", "esc":
			m.state = enums.ENDING
			return m, tea.Quit
		case "ctrl+z":
			m.suspending = true
			return m, tea.Suspend
		case "ctrl+b":
			var cmd tea.Cmd
			if m.altscreen {
				cmd = tea.ExitAltScreen
			} else {
				cmd = tea.EnterAltScreen
			}
			m.altscreen = !m.altscreen
			return m, cmd
		}
	}

	var cmd tea.Cmd
	m.list, cmd = util.ListUpdate(m.list, msg)

	return m, cmd
}

func clearScreenInAltMode(isAltScreen bool) tea.Cmd {
	if isAltScreen {
		return tea.ClearScreen
	}
	return nil
}
