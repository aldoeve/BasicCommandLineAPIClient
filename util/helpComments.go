package util

import (
	"BCLAC/styles"
)

// Returns a string about commands that alter the application run state.
func AppCommands() string {
	return styles.HelpStyle.Render("\n\n  ctrl+b: switch modes • ctrl+z: suspend • ctrl+q/esc: exit • enter: continue\n")
}

// Returns a string about keys that the user can use to navigate the application.
func NavigationHelp() string {
	return styles.HelpStyle.Render("Use the arrows to navigate • Enter/Right arrow to confirm selection • Left arrow to return\n")
}
