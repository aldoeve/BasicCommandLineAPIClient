package tui

import (
	"BCLAC/enums"
	"BCLAC/tui/views"
	"BCLAC/util"
)

func (m Model) View() string {
	if m.suspending {
		return ""
	}

	switch m.state {
	case enums.MAIN_VIEW:
		m.list.View()
		return views.FirstSelectionView(m.list) + util.NavigationHelp()
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
