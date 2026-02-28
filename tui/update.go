package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"BCLAC/enums"
	"BCLAC/tui/views"
	"BCLAC/util"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.ResumeMsg:
		m.suspending = false
		return m, nil
	case tea.WindowSizeMsg:
		m.list = util.HandleResizeofList(m.list, m.altscreen)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "right":
			m, cmd = handleForwardTravel(m)
			return m, cmd
		case "left":
			m, cmd = handleBackTravel(m)
			return m, cmd
		case "ctrl+q", "esc":
			m.state = enums.ENDING
			return m, tea.Quit
		case "ctrl+z":
			m.suspending = true
			return m, tea.Suspend
		case "ctrl+b":
			if m.altscreen {
				cmd = tea.ExitAltScreen
			} else {
				cmd = tea.EnterAltScreen
			}
			m.altscreen = !m.altscreen
			return m, cmd
		}
	}

	m.list, cmd = util.HandleListUpdates(m.list, msg)

	return m, cmd
}

func clearScreenInAltMode(isAltScreen bool) tea.Cmd {
	if isAltScreen {
		return tea.ClearScreen
	}
	return nil
}

func handleForwardTravel(m Model) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case enums.INTRO:
		m.state = enums.MAIN_VIEW
		m.list = views.InitiateFirstSelectionList(m.altscreen)
		cmd = nil
	case enums.MAIN_VIEW:
		m.state = views.ReturnNextStateFromMainSelection(uint(m.list.Index()))
		cmd = nil
	}

	return m, cmd
}

func handleBackTravel(m Model) (Model, tea.Cmd) {
	m.list.Index()
	var cmd tea.Cmd
	switch m.state {
	case enums.INTRO, enums.MAIN_VIEW:
		cmd = nil
	case enums.PASTE_FIRE, enums.QUICK_BUILD, enums.RECENTS, enums.MANUAL:
		m.state = enums.MAIN_VIEW
		m.list = views.InitiateFirstSelectionList(m.altscreen)
	}
	return m, cmd
}
