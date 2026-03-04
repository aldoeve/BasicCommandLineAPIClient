package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

// Renders the view for sending a request and the msg recived.
func HttpView(info string, spinner *spinner.Model) string {
	header := fmt.Sprintf("Checking: %s", info)
	httpRequest := fmt.Sprintf(" %s", spinner.View())
	httpResponse := ""

	view := lipgloss.JoinVertical(lipgloss.Top, header, httpRequest, httpResponse)

	return view
}
