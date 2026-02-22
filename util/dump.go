package components

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func dump(fileName string) {
	f, err := tea.LogToFile(fileName+".log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
}
