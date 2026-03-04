package tui

import (
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
			cmd = handleForwardTravel(m)
			return m, cmd
		case "left":
			cmd = handleBackTravel(m)
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

	cmd = AllowListAndTextInputUpdates(m, msg, cmd)

	return m, cmd
}

func clearScreenInAltMode(isAltScreen bool) tea.Cmd {
	if isAltScreen {
		return tea.ClearScreen
	}
	return nil
}

// Knows how to move to the next appropriate state.
func handleForwardTravel(m *Model) tea.Cmd {
	var cmd tea.Cmd
	switch m.state {
	case enums.INTRO:
		m.state = enums.MAIN_VIEW
		m.list = views.InitiateFirstSelectionList(m.altscreen)
		cmd = nil
	case enums.MAIN_VIEW:
		m.state = views.ReturnNextStateFromMainSelection(uint(m.list.Index()))
	case enums.QUICK_BUILD:
		m.state = enums.PASTE
		cmd = m.textInput.Cursor.BlinkCmd()
	case enums.PASTE:
		m.state = enums.FIRE_N_SHOW_HTTP
		cmd = tea.Batch(m.spinner.Tick)

	}

	return cmd
}

// Figures out how to move back from certain states.
func handleBackTravel(m *Model) tea.Cmd {
	var cmd tea.Cmd
	switch m.state {
	case enums.INTRO, enums.MAIN_VIEW:
		cmd = nil
	case enums.QUICK_BUILD, enums.RECENTS, enums.MANUAL:
		m.state = enums.MAIN_VIEW
		m.list = views.InitiateFirstSelectionList(m.altscreen)
		m.textInput.Reset()
	case enums.PASTE:
		m.state = enums.QUICK_BUILD

		m.textInput.Reset()
	case enums.FIRE_N_SHOW_HTTP:
		m.state = enums.PASTE
		m.textInput.Reset()
	}
	return cmd
}

// Looks at current state to mutate lists and input fields.
func AllowListAndTextInputUpdates(m *Model, msg tea.Msg, cmd tea.Cmd) tea.Cmd {
	var toBatch []tea.Cmd
	toBatch = append(toBatch, cmd)

	switch m.state {
	case enums.MAIN_VIEW:
		m.list, cmd = util.HandleListUpdates(m.list, msg)
		toBatch = append(toBatch, cmd)
	case enums.PASTE:
		m.textInput, cmd = m.textInput.Update(msg)
		toBatch = append(toBatch, cmd)
	case enums.FIRE_N_SHOW_HTTP:
		m.spinner, cmd = m.spinner.Update(msg)
		toBatch = append(toBatch, cmd)
	}

	return tea.Batch(toBatch...)
}
