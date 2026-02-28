package views

import "github.com/charmbracelet/lipgloss"

// Takes the current data in the input feild and returns the apporpriate string.
func PasteNFireView(userInput string) string {
	informUser := "Copy and Paste or Type in your http request.\n"
	str := lipgloss.JoinVertical(lipgloss.Top, informUser, userInput)
	return str
}
