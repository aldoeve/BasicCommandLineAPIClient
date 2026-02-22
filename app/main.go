package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"BCLAC/tui"
)

func main() {
	if _, err := tea.NewProgram(tui.Model{}).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
