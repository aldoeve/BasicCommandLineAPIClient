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
	case enums.PASTE_FIRE:
		return "Under Maintinence"
	case enums.QUICK_BUILD:
		return "Under Maintinence"
	case enums.RECENTS:
		return "Under Maintinence"
	case enums.MANUAL:
		return "Under Maintinence"
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
