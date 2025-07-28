package main

import (
	"fmt"
	"os"
	"s3storage/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := tui.TestModelTabs()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
