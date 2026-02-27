// Package util provides utility functions that help build and modify basic components of the application.
package util

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Dumps logging info to the given filename.
func dump(fileName string) {
	f, err := tea.LogToFile(fileName+".log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
}
