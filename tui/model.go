package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"BCLAC/tui/views"
	"BCLAC/util/enums"
)

type Model struct {
	state      enums.State
	altscreen  bool
	suspending bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.ResumeMsg:
		m.suspending = false
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
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
	return m, nil
}

func (m Model) View() string {
	if m.suspending {
		return ""
	}
	if m.state == enums.ENDING {
		return "BYE!\n"
	}

	const (
		altscreenMode = "altscreen mode"
		inlineMode    = "inline mode"
	)

	return views.IntroSequence() + HelpStyle.Render("\n\n  ctrl+b: switch modes • ctrl+z: suspend • ctrl+q: exit\n")
}
