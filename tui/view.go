package tui

import (
	"BCLAC/enums"
	"BCLAC/tui/views"
	"BCLAC/util"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	if m.suspending {
		return ""
	}

	switch m.state {
	case enums.MAIN_VIEW:
		return lipgloss.JoinVertical(lipgloss.Top, views.FirstSelectionView(m.list), util.NavigationHelp())
	case enums.QUICK_BUILD:
		return lipgloss.JoinVertical(lipgloss.Top, "Under Maintinence", util.NavigationHelp())
	case enums.PASTE:
		return lipgloss.JoinVertical(lipgloss.Top, views.PasteView(m.textInput.View()), util.NavigationHelp())
	case enums.FIRE_N_SHOW_HTTP:
		return lipgloss.JoinVertical(lipgloss.Top, views.HttpView(m.textInput.Value(), &m.spinner), util.NavigationHelp())
	case enums.RECENTS:
		return lipgloss.JoinVertical(lipgloss.Top, "Under Maintinence", util.NavigationHelp())
	case enums.MANUAL:
		return lipgloss.JoinVertical(lipgloss.Top, "Under Maintinence", util.NavigationHelp())
	case enums.ENDING:
		return "BYE!\n"
	default:
		const (
			altscreenMode = "altscreen mode"
			inlineMode    = "inline mode"
		)
		return views.IntroSequence() + util.AppCommands()
	}
}
