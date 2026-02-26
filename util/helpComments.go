package util

import (
	"BCLAC/styles"
)

func AppCommands() string {
	return styles.HelpStyle.Render("\n\n  ctrl+b: switch modes • ctrl+z: suspend • ctrl+q/esc: exit • enter: continue\n")
}

func NavigationHelp() string {
	return styles.HelpStyle.Render("\n\n Use the arrows to navigate • Enter/Right arrow to confirm selection • Left arrow to return\n")
}
