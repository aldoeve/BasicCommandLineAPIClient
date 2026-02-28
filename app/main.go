// Package main is where the application is spun up from.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"BCLAC/tui"
)

func main() {
	if _, err := tea.NewProgram(tui.InitalModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
