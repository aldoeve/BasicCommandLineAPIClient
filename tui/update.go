package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"BCLAC/enums"
	"BCLAC/tui/views"
	"BCLAC/util"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		if isUserInputNeeded(m.state) {
			m.textInput, cmd = m.textInput.Update(msg)
			cmd = tea.Batch(textinput.Blink, cmd)
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

// Knows how to move to the next appropriate state.
func handleForwardTravel(m *Model) (*Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case enums.INTRO:
		m.state = enums.MAIN_VIEW
		m.list = views.InitiateFirstSelectionList(m.altscreen)
		cmd = nil
	case enums.MAIN_VIEW:
		m.state = views.ReturnNextStateFromMainSelection(uint(m.list.Index()))
		if m.state == enums.PASTE_FIRE {
			cmd = textinput.Blink
		}
	}

	return m, cmd
}

// Figures out how to move back from certain states.
func handleBackTravel(m *Model) (*Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case enums.INTRO, enums.MAIN_VIEW:
		cmd = nil
	case enums.PASTE_FIRE, enums.QUICK_BUILD, enums.RECENTS, enums.MANUAL:
		m.state = enums.MAIN_VIEW
		m.list = views.InitiateFirstSelectionList(m.altscreen)
		m.textInput.Reset()
	}
	return m, cmd
}

func isUserInputNeeded(state enums.State) bool {
	return state == (enums.PASTE_FIRE)
}
